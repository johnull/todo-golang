FROM golang:alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/webapp/main.go

CMD ["./main"]

