# setup base image
FROM golang:1.17.0-stretch

WORKDIR /app

COPY ./ /app

COPY ./newdb/01.init.sh /docker-entrypoint-initdb.d/

RUN go get github.com/go-redis/redis

RUN go mod tidy

ENTRYPOINT [ "go", "run", "main.go" ]