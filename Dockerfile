# ---------- build stage ----------
FROM golang:1.23-alpine AS builder
WORKDIR /app

# go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# остальной код
COPY . .

# собираем бинарь (ключевой путь!)
RUN CGO_ENABLED=0 GOOS=linux go build -o bookstore ./cmd/app

# ---------- run stage ----------
FROM alpine:3.20
WORKDIR /app

COPY --from=builder /app/bookstore .
COPY --from=builder /app/static ./static

ENV PORT=8080 APP_ENV=production
EXPOSE 8080
ENTRYPOINT ["./bookstore"]
