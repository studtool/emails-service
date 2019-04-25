#!/usr/bin/env bash

command="$1"

repository=$(cut "-d " "-f2" ./go.mod | head -n 1)

app=$(echo ${repository} | cut "-d/" "-f2")
service=$(echo ${repository} | cut "-d/" "-f3")
version=$(cut "-d " "-f4" ./go.mod | head -n 1)

image="${app}/${service}:${version}"

if [[ "${command}" = "build" ]]; then
  docker build -t "${image}" .
elif [[ "$command" = "push" ]]; then
  echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin \
    && docker push "$image"
elif [[ "${command}" = "remove" ]]; then
  docker rmi "${image}"
else
  echo "command expected 'build/push/remove'"
  exit -1
fi
