# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod files and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy the whole source code
COPY ./src ./src

# Build
WORKDIR /app/src/cmd
RUN go build -o /innoscripta-cs

# Final stage
FROM alpine:3.18

COPY --from=builder /innoscripta-cs /innoscripta-cs

EXPOSE 8080

ENTRYPOINT ["/innoscripta-cs"]
