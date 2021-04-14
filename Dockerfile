FROM golang:1.14-alpine3.11 as baseimage

RUN mkdir /app
FROM baseimage
WORKDIR /app
COPY . .
RUN go mod download
WORKDIR /app/game
ENV CGO_ENABLED=0
RUN go build
RUN go test