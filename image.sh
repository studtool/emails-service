#!/usr/bin/env bash

su=""
if [[ $(id -u) = 0 ]]; then
    su="sudo"
fi

command="$1"

args=""
for (( i=2; i<=$#; i++ )) do
    args+="${!i} "
done

repository=$(cut "-d " "-f2" ./go.mod | head -n 1)

app=$(echo ${repository} | cut "-d/" "-f2")
service=$(echo ${repository} | cut "-d/" "-f3")
version=$(sed '3q;d' ./go.mod | cut "-d " "-f2")

image="${app}/${service}:${version}"

if [[ "${command}" = "build" ]]; then
    shcmd="docker build -t ${image} ."
elif [[ "${command}" = "run" ]]; then
    shcmd="docker run ${args} ${image}"
elif [[ "$command" = "push" ]]; then
    echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin \
        && shcmd="docker push ${image}"
elif [[ "${command}" = "remove" ]]; then
    shcmd="docker rmi ${image}"
else
  echo "command expected 'build/push/remove'"
  exit -1
fi

echo "${su} ${shcmd}" && sh -c "${shcmd}"
