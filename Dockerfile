FROM golang:1.25.4 AS builder
WORKDIR /app

RUN go env -w CGO_ENABLED=0 GOOS=linux GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/user-service ./cmd/users-service

FROM alpine:latest
WORKDIR /service
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/user-service .
RUN chmod +x user-service
EXPOSE 8080
CMD ["./user-service"]
