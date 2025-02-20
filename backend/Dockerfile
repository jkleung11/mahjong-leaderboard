# build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app

RUN apk add --no-cache gcc musl-dev sqlite sqlite-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# runtime stage
FROM alpine:latest

WORKDIR /root/
RUN apk add --no-cache sqlite

COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]