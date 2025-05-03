# Build stage
FROM golang:1.19-alpine AS builder
WORKDIR /app

# First copy just the dependencies files
COPY go.mod go.sum ./
RUN go mod download

# Then copy the rest of the application
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o match-me .

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/match-me .
COPY --from=builder /app/.env .  # Make sure this exists
EXPOSE 8080
CMD ["./match-me"]