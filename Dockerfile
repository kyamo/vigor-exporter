FROM golang AS build

LABEL maintainer="Kyamo" \
      version="latest" \
      description="Prometheus exporter for Vigor modems"

COPY . /go/src/github.com/kyamo/vigor-exporter

WORKDIR /go/src/github.com/kyamo/vigor-exporter

RUN apt update && \
	apt upgrade -y && \
	apt install curl  -y

ENV GOPATH=/go
RUN go build -tags netgo -a -v -ldflags "-s -w" main.go 

# Final image
FROM debian:11-slim

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get clean -y

COPY --from=build /go/src/github.com/kyamo/vigor-exporter /vigor-exporter

ENV HOST=192.168.0.1 \
    USERNAME=admin \
	PASSWORD=changeme

VOLUME /vigor-exporter

WORKDIR /vigor-exporter

EXPOSE 9103

ENTRYPOINT [ "/bin/sh", "-c", "./main --host $HOST --password $PASSWORD --username $USERNAME" ]
