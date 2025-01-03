# Кредитный калькулятор

Этот проект представляет собой веб-приложение для расчета общей суммы кредита и ежемесячного платежа.

## Функционал

- Ввод суммы кредита, процентной ставки и срока кредита.
- Расчет общей суммы к оплате.
- Расчет ежемесячного платежа.
- Удобный интерфейс с валидированием данных.

## Стек технологий

- **Go (Golang)** — основная логика приложения.
- **Chi** — маршрутизатор для обработки HTTP-запросов.
- **HTML и CSS** — фронтенд, отображающий форму и результаты.

## Установка и запуск

1. Склонируйте репозиторий:

    ```bash
    git clone https://github.com/ваш-репозиторий.git
    cd ваш-репозиторий
    ```

2. Убедитесь, что у вас установлен Go.

3. Установите зависимости:

    ```bash
    go mod tidy
    ```

4. Запустите сервер:

    ```bash
    go run main.go
    ```

5. Перейдите в браузере по адресу [http://localhost:8080](http://localhost:8080).

## Тестирование

Для запуска тестов выполните:

```bash
go test ./...
```
