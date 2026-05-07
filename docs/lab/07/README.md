# Лабораторная работа №7

## Подготовка

⚠️ **Обязательно ознакомьтесь с
[CONTRIBUTING.md](https://github.com/sitnikovik/ndbx-template?tab=contributing-ov-file)** -
там описан процесс работы с репозиторием, Pull Requests и GitHub Actions.

## Цель работы

Реализовать функционал рекомендаций мероприятий c использованием [Neo4j](https://neo4j.com/)

> Интересы пользователей хранятся в графовой базе данных Neo4j как связи между пользователем и тегами мероприятий.
> Это позволяет строить рекомендации мероприятий на основе схожих интересов пользователей.

## Эндпоинты

### Интересы пользователя

Реализуйте новый эндпоинт `PUT /users/{user_id}/interests`, который принимает список интересов пользователя
и сохраняет их в Neo4j для построения рекомендаций мероприятий.

> 🔐 Доступно **только авторизованным** пользователям. Пользователь может менять **только свои** интересы.

**Запрос**:

```http
PUT /users/{user_id}/interests HTTP/1.1
Host: localhost:8080
Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d;
Content-Type: application/json
Content-Length: 999
```

> ⚡️ PUT запрос полностью обновляет данные несмотря на то были ли они уже заданы или инет

**Тело запроса**:

```json
{
    "interests": ["concert", "exhibition"]
}
```

Поле `interests` обязательно и может содержать только эти значения-константы:

- concert
- exhibition
- culture
- theater
- sport
- food
- education
- technology

**Ответ (обновлены или добавлены впервые):**

```http
HTTP/1.1 204 No Content
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/ Max-Age={APP_USER_SESSION_TTL}
Content-Length: 0
```

**Ответ (пустой массив интересов):**

```http
HTTP/1.1 400 Bad Request
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/ Max-Age={APP_USER_SESSION_TTL}
Content-Length: 999
Content-Type: application/json
{"message": "field \"interests\" is empty"}
```

**Ответ (некорректное значение в интересах):**

```http
HTTP/1.1 400 Bad Request
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/ Max-Age={APP_USER_SESSION_TTL}
Content-Length: 999
Content-Type: application/json
{"message": "field \"interests\" contains incorrect value \"cinema\""}
```

**Ответ (пользователь не найден или попытка поменять интересы другому пользователю):**

```http
HTTP/1.1 404 Not Found
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/ Max-Age={APP_USER_SESSION_TTL}
Content-Type: application/json
Content-Length: 28
{"message": "User not found"}
```

> ⚠️ Пользователь может менять только свои интересы

**Ответ (если пользователь не авторизован):**

```http
HTTP/1.1 401 Unauthorized
Content-Length: 0
```

### Теги мероприятий

Добавьте возможность добавлять/изменять теги для мероприятия
в существующие [PATCH-запрос](/docs/lab/04/README.md#редактирование-мероприятий) и [POST-запрос](/docs/lab/03/README.md#создание-события)

В теле запроса добавляется поле `tags` со значениями [интересов](#интересы):

```json
{
    "tags": ["concert", "culture", "theater"]
}
```

**Ответ (если передан пустой список тегов):**

```http
HTTP/1.1 400 Bad Request
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/ Max-Age={APP_USER_SESSION_TTL}
Content-Type: application/json
Content-Length: 999
{"message": "field \"tags\" is empty"}
```

> ⚠️ Все эндпоинты получения мероприятиям должны возвращать
поле `tags` с заданным списком тегов.

### Рекомендации мероприятий

Возвращает список рекомендованных мероприятий на основе интересов пользователя и графа в Neo4j.

> 🔐 Только для авторизованных пользователей. Пользователь может запрашивать рекомендации только для себя.

```http
GET /users/{user_id}/recommendations HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Content-Length: 999
```

**Ответ (ОК):**

```http
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: 999
```

```json
{
  "recommendations": [
    {
      "id": "event_123",
      "title": "Международная выставка современного искусства",
      "date": "2025-04-15T18:00:00Z",
      "tags": ["exhibition", "culture"],
      "score": 0.92
    },
    {
      "id": "event_456",
      "title": "Концерт симфонического оркестра",
      "date": "2025-04-17T19:30:00Z",
      "tags": ["concert", "culture"],
      "score": 0.85
    }
  ]
}
```

## Интересы

Список значений-констант:

- `concert`
- `exhibition`
- `culture`
- `theater`
- `sport`
- `food`
- `education`
- `technology`

## Схема данных в Neo4j

**Узлы:**

- `User` — пользователь системы
  - `id` *string* — его идентификатор в MongoDB
- `Interest` — тематический интерес
  - `name` — одно из интересов (см. выше)
- `Event` — мероприятие
  - `id` *string* — его идентификатор в MongoDB
  - `title` *string* — название

**Связи:**

- `(u:User)-[:HAS_INTEREST]->(i:Interest)` — пользователь заинтересован в теме
- `(e:Event)-[:HAS_TAG]->(i:Interest)` — мероприятие относится к теме

**Индексы:**

- `FOR (u:User) ON (u.id)` — для быстрого поиска пользователя
- `FOR (i:Interest) ON (i.name)` — для быстрого сопоставления интересов

## Конфигурация

Добавьте параметры конфигурации контейнера с Neo4j в `.env.local`

```sh
NEO4J_HOST=neo4j
NEO4J_USER=neo4j
NEO4J_PASSWORD=password
```

## FAQ

**Q: Нужен ли эндпоинт для получения интересов?**  
A: Нет, только добавляем/переписываем через PUT-запрос.

**Q: Нужно ли возвращать список доступных значений интересов в сообщение при ошибке в PUT-запросе?**  
A: В нашем случае — не обязательно.

**Q: Почему интересы как отдельный узел, а не свойство?**  
A: Интересы вынесены в отдельные узлы, чтобы строить
графовые связи между пользователями и мероприятиями через общие темы
