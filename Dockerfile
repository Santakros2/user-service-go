FROM oraclelinux:8 AS builder

# Install instant client
RUN dnf install -y oracle-instantclient-release-el8 \
    && dnf install -y oracle-instantclient-basic oracle-instantclient-devel

RUN dnf install -y gcc gcc-c++ make

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o users-service ./cmd/users-service

FROM oraclelinux:8

RUN dnf install -y oracle-instantclient-basic

WORKDIR /app
COPY --from=builder /app/users-service .

EXPOSE 8080
CMD ["./users-service"]
