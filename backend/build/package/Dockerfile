FROM golang:1.24-alpine AS builder

WORKDIR /app

# Копируем файлы зависимостей первыми для лучшего кэширования
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение с оптимизациями
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# Финальная стадия
FROM alpine:3.19

WORKDIR /app

# Копируем бинарник из стадии сборки
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Запускаем приложение
CMD [ "./main" ]