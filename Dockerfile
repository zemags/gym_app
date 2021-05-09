FROM golang:1.15-alpine as builder

WORKDIR /opt/code/
ADD ./ /opt/code

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN go mod download

RUN go build -o bin/gymapp_binary cmd/main.go
ENTRYPOINT ["bin/gymapp_binary"]