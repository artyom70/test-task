# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

WORKDIR /go/src/app
COPY . .

RUN go mod download

EXPOSE 8080

RUN go build -o /go/bin/app -mod vendor ./main.go


CMD [ "/go/bin/app" ]
