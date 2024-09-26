FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY config/ ./config/
COPY internal/ ./internal/
COPY migrations/ ./migrations/
COPY pkg/ ./pkg/

RUN go build -o serenade ./cmd/app/main.go

# Use a larger base image that includes PostgreSQL
FROM postgres:latest

# Set environment variables for PostgreSQL
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=yourpassword
ENV POSTGRES_DB=mydatabase

# Copy built Go application into the PostgreSQL container
COPY --from=builder /app/serenade /usr/local/bin/serenade
COPY --from=builder /app/migrations/ /usr/local/bin/migrations

# Expose the ports for both services (if needed)
EXPOSE 5432
EXPOSE 8080

# Command to start both PostgreSQL and the Go app (this is not standard practice)
CMD ["sh", "-c", "serenade"]
