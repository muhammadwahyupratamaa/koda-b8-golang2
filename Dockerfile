FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o koda-b8-golang2 main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/koda-b8-golang2 .

CMD ["./koda-b8-golang2"]