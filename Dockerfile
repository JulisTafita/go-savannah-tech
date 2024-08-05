FROM ubuntu:latest
LABEL authors="Julis"

FROM golang:1.22

WORKDIR /app
ADD . /app/

RUN go build -o ./out/go-savannahTech .

EXPOSE 8089

ENTRYPOINT ["./out/go-savannahTech"]