# syntax=docker/dockerfile:1
FROM golang:1.17-alpine
WORKDIR /u01
RUN mkdir -p /u01/data/sd_configs/redis
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /redis-prometheus
EXPOSE 9091
CMD [ "/redis-prometheus" ]