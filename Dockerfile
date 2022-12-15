FROM golang:1.18-bullseye AS builder

WORKDIR /app

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .

RUN apt-get update && \
    apt-get install -y build-essential && \
    go build -o ./server ./cmd/server

# -=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
FROM debian:bullseye

WORKDIR /app

RUN mkdir /app/config && \
    apt-get update && \
    apt-get install -y ca-certificates

COPY --from=builder /app/config ./config
COPY --from=builder /app/server .

CMD ["/app/server"]
