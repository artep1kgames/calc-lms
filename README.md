# Веб-сервис для вычисления арифметических выражений

## Описание проекта
Проект представляет собой веб-сервис для выполнения арифметических вычислений. Пользователь отправляет HTTP POST-запрос с арифметическим выражением, а в ответ получает результат вычисления или сообщение об ошибке.

---

## API Документация

### **Эндпоинт**
**POST** `/api/v1/calculate`

### Формат запроса
Тело запроса должно быть в формате JSON:
```json
{
    "expression": "ваше выражение"
}
```
- **expression**: строка с арифметическим выражением для вычисления.

### Формат ответов
- **Успех (200):**
    ```json
    {
        "result": "результат вычисления"
    }
    ```
- **Ошибка 422 (Некорректное выражение):**
    ```json
    {
        "error": "Expression is not valid"
    }
    ```
- **Ошибка 500 (Серверная ошибка):**
    ```json
    {
        "error": "Internal server error"
    }
    ```

---

## Примеры использования

### Успешный запрос
```bash
curl -X POST http://localhost:8080/api/v1/calculate \
    -H "Content-Type: application/json" \
    -d '{"expression": "2+2*2"}'
```
**Ответ:**
```json
{
    "result": 6.000000
}
```

### Ошибка 422 (Некорректное выражение)
```bash
curl -X POST http://localhost:8080/api/v1/calculate \
    -H "Content-Type: application/json" \
    -d '{"expression": "2+2*2a"}'
```
**Ответ:**
```json
{
    "error": "Expression is not valid"
}
```

### Ошибка 500 (Деление на ноль)
```bash
curl -X POST http://localhost:8080/api/v1/calculate \
    -H "Content-Type: application/json" \
    -d '{"expression": "1/0"}'
```
**Ответ:**
```json
{
    "error": "Internal server error"
}
```

---

## Инструкция по запуску

### Подготовка
1. Склонируйте репозиторий:
    ```bash
    git clone https://github.com/Portalshik/calc-LMS.git
    cd calc-LMS
    ```

2. Убедитесь, что у вас установлена версия Go 1.18 или выше.

### Запуск
- Для Linux и Windows используйте команду:
    ```bash
    go run cmd/calc/main.go
    ```

Сервер будет доступен по адресу: `http://localhost:8080`.

---

## Требования
- Go версии 1.18 или выше.
- HTTP-клиент (например, `curl`).

---

## Структура проекта
```
calc-LMS/
├── cmd/
│   └── calc/
│       └── main.go
├── go.mod
├── internal/
│   ├── api/v1/
│   │   └── api.go
│   ├── calculator/
│   │   └── calculator.go
│   └── server/
│       └── server.go
└── README.md
