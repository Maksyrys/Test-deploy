FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bookstore ./cmd/app

FROM alpine:3.20
WORKDIR /app

COPY --from=builder /app/bookstore .
COPY --from=builder /app/static    ./static
COPY --from=builder /app/templates ./templates

ENV PORT=8080 APP_ENV=production
EXPOSE 8080
CMD ["./bookstore"]