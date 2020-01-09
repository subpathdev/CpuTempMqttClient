FROM golang:1.12.9-alpine3.10 AS builder
COPY . /go/src/github.com/subpathdev/CpuTempMqttClient
WORKDIR /go/src/github.com/subpathdev/CpuTempMqttClient
RUN GO111MODULE=on GOPROXY=https://proxy.golang.org CGO_ENABLED=0 go build -v -o /usr/local/bin/CpuTempMqttClient -ldflags="-w -s" github.com/subpathdev/CpuTempMqttClient

FROM ubuntu:18.04
COPY --from=builder /usr/local/bin/CpuTempMqttClient /CpuTempMqttClient
RUN apt-get update && apt-get install -y \
    lm-sensors \
 && rm -rf /var/lib/apt/lists/*
ENTRYPOINT ["/CpuTempMqttClient"]
