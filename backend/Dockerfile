# syntax=docker/dockerfile:1

FROM golang:1.18 as builder
# Download gcc required by sqlite3
RUN apt update && apt install build-essential -y
WORKDIR /app
# Download Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
RUN echo $PWD
RUN go build -o ./bin ./cmd/server

FROM ubuntu:20.04
WORKDIR /
COPY --from=builder /app/bin ./app
EXPOSE 8080
RUN groupadd -r nonroot -g 901 && useradd -u 901 -r -g nonroot nonroot
RUN chmod 755 ./app
USER nonroot:nonroot
ENTRYPOINT ["sh", "/app"]
