# Build stage
FROM golang:1.19-alpine AS builder
WORKDIR /app

# Copy module files first
COPY go.mod go.sum ./
RUN go mod download all  # Changed to download all dependencies

# Copy everything else
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o match-me .

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/match-me .
EXPOSE 8080
CMD ["./match-me"]