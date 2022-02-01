# simple docker configuration file to build container
FROM golang:1.16-alpine

WORKDIR /app

COPY *.go ./

RUN go build -o /docker-hello

CMD ["go","run","test.go"]