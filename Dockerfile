FROM golang:1.12-alpine3.9 as base
WORKDIR /tmp/emails-service
COPY . .
RUN go build -mod vendor -o /tmp/service .

FROM ubuntu:18.04
RUN apt-get update \
    && apt-get install -y postfix
WORKDIR /tmp
COPY --from=base /tmp/service ./service
ENTRYPOINT ["./service"]
EXPOSE 80
