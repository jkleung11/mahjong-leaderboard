FROM golang:1.23

WORKDIR /app

RUN apt-get update && apt-get install -y sqlite3
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]