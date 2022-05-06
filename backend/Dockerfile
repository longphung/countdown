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
RUN go build -o ./bin ./cmd/server

FROM ubuntu:20.04
ENV GIN_MODE=release
ENV DB_LOCATION=/home/nonroot/db/db.sqlite
RUN groupadd -r nonroot -g 901 && useradd -u 901 -r -g nonroot -m nonroot
WORKDIR /home/nonroot
RUN mkdir db
VOLUME /home/nonroot/db
COPY --from=builder /app/bin /home/nonroot/app
EXPOSE 8080
RUN chmod 755 /home/nonroot/app && chown nonroot:nonroot /home/nonroot/app
RUN chmod 755 /home/nonroot/db && chown nonroot:nonroot /home/nonroot/db
USER nonroot:nonroot
ENTRYPOINT ["./app"]
