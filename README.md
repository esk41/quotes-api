# Quotes API

Простой REST API-сервис для хранения и управления цитатами на Go.

## Запуск

```bash
go mod tidy
go run .
```

## Примеры запросов

```bash
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'

curl http://localhost:8080/quotes
curl http://localhost:8080/quotes/random
curl http://localhost:8080/quotes?author=Confucius
curl -X DELETE http://localhost:8080/quotes/1
```
## Особенности

- Хранение цитат в памяти
- Фильтрация по автору
- Выдача случайной цитаты
- Отслеживание времени добавления (`created_at`)

### Получение цитаты по ID

**GET /quotes/{id}**

Возвращает цитату по заданному ID.

**Пример запроса:**

```bash
curl http://localhost:8080/quotes/1