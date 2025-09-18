# ServiceBuyer API

Сервис для управления подписками пользователей с Swagger документацией.

## Описание

ServiceBuyer - это REST API сервис, который позволяет:
- Создавать записи о подписках пользователей
- Получать статистику по подпискам с возможностью фильтрации

## Установка и запуск

1. Установите зависимости:
```bash
go mod tidy
```

2. Создайте файл `.env` с переменными окружения:
```
POSTGRES_CONNECTION_STRING="user=postgres password=password dbname=projectDB sslmode=disable"
SERVICE_SERVER_PORT=8080
```

3. Запустите сервер:
```bash
go run cmd/main.go
```

## Swagger документация

После запуска сервера Swagger UI будет доступен по адресу:
- http://localhost:8080/swagger/

## API Endpoints

### POST /record
Создает новую запись о подписке пользователя.

**Тело запроса:**
```json
{
  "name": "Netflix",
  "cost": 9.99,
  "user_id": "123e4567-e89b-12d3-a456-426614174000",
  "start_date": "2024-01-01",
  "end_date": "2024-12-31"
}
```

**Ответ:**
```json
{
  "status": "success"
}
```

### GET /count
Получает статистику по подпискам с возможностью фильтрации.

**Тело запроса:**
```json
{
  "start_date": "2024-01-01",
  "end_date": "2024-12-31",
  "user_id": "123e4567-e89b-12d3-a456-426614174000",
  "service_name": "Netflix"
}
```

**Ответ:**
```json
{
  "status": "success",
  "data": {
    "user_id": "123e4567-e89b-12d3-a456-426614174000",
    "amount": 19.98,
    "services": [
      {
        "id": 1,
        "name": "Netflix",
        "amount": 9.99
      }
    ],
    "start_date": "2024-01-01",
    "end_date": "2024-12-31"
  }
}
```

## Структура проекта

```
ServiceBuyer/
├── cmd/
│   └── main.go                 # Точка входа приложения
├── internal/
│   ├── config/                 # Конфигурация
│   ├── handler/                # HTTP обработчики
│   ├── model/                  # Модели данных
│   ├── repository/             # Работа с базой данных
│   ├── service/                # Бизнес-логика
│   └── transport/              # HTTP транспорт
├── docs/                       # Swagger документация
└── README.md
```

## Технологии

- Go 1.24
- PostgreSQL
- Gorilla Mux
- Swagger/OpenAPI 3.0
- Logrus для логирования
