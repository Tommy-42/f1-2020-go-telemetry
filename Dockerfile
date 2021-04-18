FROM golang:1.16.3-alpine AS builder
WORKDIR /f1
COPY . .
RUN go build -o /f1/bin/f1-2020-ingester .

ENTRYPOINT ["/f1/bin/f1-2020-ingester"]
