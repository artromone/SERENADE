FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY pkg/ ./pkg/
COPY main/ ./main/

RUN go build -o serenade ./cmd/app/main.go

EXPOSE 8080
CMD ["./serenade"]
