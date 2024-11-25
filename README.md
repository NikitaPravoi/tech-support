# tech-support

## Описание

Данный репозиторий реализует функционал проектов, создания и ответа на запросы в техподдержку. Авторизация доступна на основе сессий.

## Документация по проекту

**[docs](docs/docs.yaml)**

**[huma](https://huma.rocks/)**

**[Структура проекта](https://habr.com/ru/companies/inDrive/articles/690088/)**


## Кратко по директориям проекта

### `/cmd`

Исходная точка для запуска проекта

### `/config`

Место для хранения файлов конфигурации

### `/docs`

Место для хранения документации

### `/internal`

Место для хранения кодовой базы проекта

### `/sql`

Место для хранения схемы и запросов SQL, необходимо для работы sqlc

## TODO:

1. [ ] Добавить конфиг
2. [ ] Добавить метод PATCH на ticket для добавления ответа
3. [ ] Переделать RolesMiddleware