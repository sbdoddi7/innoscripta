FROM golang:1.21-alpine AS builder

WORKDIR /app

# copy go.mod & download deps first for cache
COPY go.mod go.sum ./
RUN go mod download

# copy source
COPY ./src ./src

# build
RUN cd src && go build -o /innoscripta-cs

# Final stage
FROM alpine:3.18

# copy binary
COPY --from=builder /innoscripta-cs /innoscripta-cs

EXPOSE 8080

ENTRYPOINT ["/innoscripta-cs"]