# Start from golang base image
FROM golang:alpine as builder

# Add a Maintainer inf
LABEL maintaner="Snow Dev <inoue.sora@gmail.com>"

WORKDIR /home/titor/Documents/go/src/github.com/snow-dev/super_blog

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./main.go" -command="./main"