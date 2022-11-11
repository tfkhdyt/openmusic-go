FROM golang:alpine3.16

WORKDIR /app

COPY . .

RUN go build -o bin/openmusic-go main.go

EXPOSE 8080

ENV GIN_MODE=release

CMD bin/openmusic-go
