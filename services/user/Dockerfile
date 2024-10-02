FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -a -o main main.go

FROM ubuntu:24.04
WORKDIR /app
COPY --from=builder /app/main /app
COPY .env.example .env
CMD ["./main"]