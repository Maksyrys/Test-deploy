# ---------- build stage ----------
FROM golang:1.22-alpine AS builder
WORKDIR /app

# кешируем зависимости
COPY src/go.mod src/go.sum ./
RUN go mod download

# копируем всё остальное
COPY src/ .

# соберём статически (без CGO) под Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o bookstore ./cmd/bookstore

# ---------- run stage ----------
FROM alpine:3.20
WORKDIR /app

# бинарь + статику
COPY --from=builder /app/bookstore .
COPY --from=builder /app/static ./static

# переменные окружения по-умолчанию (можно переопределять в compose/.env)
ENV PORT=8080 \
    APP_ENV=production

EXPOSE 8080
CMD ["./bookstore"]