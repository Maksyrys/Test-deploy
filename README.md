# BookStore — Интернет-магазин книг на Go

BookStore — полнофункциональный веб-приложение для покупки книг, реализованное на языке Go с использованием чистой архитектуры, PostgreSQL, Docker и шаблонов HTML.  
Проект разработан в рамках дипломной работы.

## Возможности

- Регистрация и аутентификация пользователей (сессии)
- Просмотр каталога книг, поиск, фильтрация по категориям и авторам
- Страница детальной информации о книге с отзывами пользователей
- Добавление книг в корзину, оформление заказа
- Избранное
- Личный кабинет пользователя с редактированием профиля и историей отзывов
- Админ-панель для управления книгами и отзывами
- Современный UI (HTML, CSS, JS), поддержка адаптивности

## Технологический стек

- **Go 1.23** (net/http, gorilla/mux, gorilla/sessions, zap logger)
- **PostgreSQL** (16+)
- **Docker** и docker-compose
- **Nginx** (reverse proxy, статика)
- **HTML-шаблоны**
- **Migrations** (migrate/migrate)
- **bcrypt** для паролей

## Быстрый старт (Docker)

1. **Клонируйте репозиторий:**
    ```sh
    git clone https://github.com/Maksyrys/Test-deploy.git
    cd Test-deploy
    ```

2. **Настройте переменные окружения:**
    > Можно оставить значения по умолчанию или скорректировать в файле `.env` (создайте на основе `.env.example`).

3. **Запустите проект:**
    ```sh
    docker-compose up --build
    ```
    - По умолчанию приложение будет доступно на [http://localhost](http://localhost)
    - Статические файлы доступны по адресу `/static/`

4. **Примените миграции (опционально, если migrations не применились):**
    ```sh
    docker-compose run --rm migrate
    ```

5. **Остановка проекта:**
    ```sh
    docker-compose down
    ```

## Структура проекта

- `/cmd/app` — точка входа приложения
- `/internal/app` — слои приложения (handlers, services, middlewares)
- `/internal/models` — структуры данных
- `/internal/repository` — интерфейсы и реализации для работы с БД
- `/internal/config` — работа с переменными окружения и настройками
- `/migrations` — SQL-миграции для PostgreSQL
- `/templates` — HTML-шаблоны (layout, index, book, cart и др.)
- `/static` — CSS, JS, картинки

## Основные команды (локальная разработка)

- **Сборка приложения:**  
  `go build -o bookstore ./cmd/app`
- **Запуск (без Docker):**  
  Настройте переменные среды (`BD_HOST`, `BD_PORT`, и т.д.), затем:  
  `./bookstore`
- **Миграции:**  
  Используйте контейнер или локально инструмент [migrate](https://github.com/golang-migrate/migrate):

  ```sh
  migrate -path ./migrations -database "postgres://user:password@localhost:5432/bookstore?sslmode=disable" up
