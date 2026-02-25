# Stage 1: Сборка
FROM golang:1.25-alpine AS build
WORKDIR /app
# Сначала копируем файлы модулей для кэширования слоев
COPY go.mod go.sum ./
RUN go mod download
# Копируем остальное и собираем
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/app ./cmd/main.go

# Stage 2: Финальный образ
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /app/bin/app .
# Копируем SQL-скрипты, если они нужны при старте (опционально)
# COPY --from=build /app/migrations ./migrations 

EXPOSE 8080
CMD ["./app"]
