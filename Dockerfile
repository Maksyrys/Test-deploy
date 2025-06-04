# ---------- build stage ----------
FROM golang:1.22-alpine AS builder
WORKDIR /app

# 1. go.mod + go.sum
COPY go.mod go.sum ./
RUN go mod download

# 2. всё остальное
COPY . .

# 3. собираем статически
RUN CGO_ENABLED=0 GOOS=linux go build -o bookstore ./cmd/bookstore

# ---------- run stage ----------
FROM alpine:3.20
WORKDIR /app

COPY --from=builder /app/bookstore .
COPY --from=builder /app/static ./static   # если есть статические файлы

ENV PORT=8080 APP_ENV=production
EXPOSE 8080
ENTRYPOINT ["./bookstore"]
