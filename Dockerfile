FROM golang:1.17

WORKDIR /go/src/app

COPY HttpServer.go /HttpServer.go

EXPOSE 8888

ENTRYPOINT go run /HttpServer.go