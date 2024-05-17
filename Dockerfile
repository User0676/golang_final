FROM golang:1.21

# Устанавливаем переменные среды для подключения к базе данных PostgreSQL
ENV DB_HOST=localhost \
    DB_PORT=5432 \
    DB_USER=postgres \
    DB_NAME=postgres \
    DB_SSLMODE=disable

# Создаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем все остальные файлы проекта
COPY . .

# Компилируем и собираем приложение
RUN go build -o gym-app main.go

# Команда для запуска приложения при старте контейнера
CMD ["./gym-app"]