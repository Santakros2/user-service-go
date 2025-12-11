# ------------ Stage 1: Build ---------------------
FROM golang:1.25.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o users-service ./cmd/users-service

# ------------ Stage 2: Run -----------------------
FROM ubuntu:22.04

WORKDIR /app

COPY --from=builder /app/users-service /app/users-service
COPY .env /app/.env

RUN apt-get update && apt-get install -y ca-certificates

EXPOSE 8080

CMD ["/app/users-service"]
