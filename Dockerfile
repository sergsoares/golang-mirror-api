FROM golang:1.18.3 AS builder

COPY ./mirror /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o program

ENTRYPOINT [ "./program" ]