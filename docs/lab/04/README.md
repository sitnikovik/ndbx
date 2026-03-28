# Лабораторная работа №4

## Подготовка

⚠️ **Обязательно ознакомьтесь с
[CONTRIBUTING.md](https://github.com/sitnikovik/ndbx-template?tab=contributing-ov-file)** -
там описан процесс работы с репозиторием, Pull Requests и GitHub Actions.

## Цель работы

Реализовать шардирование и репликацию в MongoDB для коллекции `events` и `users`,
а также добавить новые эндпоинты для редактирования мероприятий, поиска мероприятий и организаторов.

## Задание

### Редактирование мероприятий

Реализуйте новый эндпоинт `PATCH /events/{id}`, позволяющий редактировать данные о мероприятии.

> 🔐 Доступ только у организатора мероприятия

```http
PATCH /events/{id} HTTP/1.1
Host: localhost:8000
Content-Type: application/json
Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d;
```

В теле запроса добавляется несколько полей:

- Категория
- Цена билета
- Наименование города

```json
{
    "category": "концерт",
    "price": 1000,
    "city": "Москва"
}
```

> ⚡️ Все цены передаем не в валюте, а просто в виде абстрактной единицы и только **целые числа**!

- `category` *string* - категория мероприятия, добавляет/меняет в events.category в MongoDB,
может принимать следующие значения:
  - `meetup`
  - `concert`
  - `exhibition`
  - `party`
  - `other`
- `price` *uint* - цена билета, добавляет/меняет в events.price в MongoDB
- `city` *string*:
  - при наличии значения добавляет/меняет в `events.location.city` в MongoDB
  - при пустом значении в теле запроса - удаляет `events.location.city` в MongoDB

**Ответ (при успешном обновлении):**

```http
HTTP/1.1 204 No Content
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Length: 0
```

**Ответ (если мероприятие не найдено или пользователь не является его организатором):**

```http
HTTP/1.1 404 Not Found
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Length: 0
```

**Ответ (если какой-то из параметров невалиден):**

```http
HTTP/1.1 400 Bad Request
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Type: application/json
Content-Length: 999
{"message": "invalid \"{field_name}\" field"}
```

**Ответ (если пользователь не авторизован):**

```http
HTTP/1.1 401 Unauthorized
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Length: 0
```

### Поиск мероприятий

Доработайте существующий эндпоинт `GET /events`, возвращающий список мероприятий, отвечающий параметрам поиска.

```http
GET /events?category=party&city=Москва&price_to=0&started_date_from=20260314 HTTP/1.1
Host: localhost:8000
Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d
```

Параметры фильтрации задаются через **GET-параметры**:

- Все параметры из [предыдущей лабораторной работы](../03/README.md#просмотр-событий)
- `id` *string* - id мероприятия (поиск по точному совпадению)
- `category` *string* - категория мероприятия
  - `meetup`
  - `concert`
  - `exhibition`
  - `party`
  - `other`
- `price_from` *uint* - минимальная цена билета включительно
- `price_to` *uint* - максимальная цена билета включительно
- `city` *string* - город проведения мероприятия
- `date_from` *string* - дата начала мероприятия не раньше этого значения (формат YYYYMMDD, например 20260314)
- `date_to` *string* - дата начала мероприятия не позже этого значения (формат YYYYMMDD, например 20260314)
- `user` *string* - никнейм пользователя, который создал мероприятие (поиск по точному совпадению)

> 💡 Чтобы отобразить только бесплатные встречи - достаточно передать `price_to=0`

**Ответ (события найдены):**

```http
HTTP/1.1 200 ОК
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Type: application/json
Content-Length: 759
```

```json
{
    "events": [
        {
            "id": "12e9c0b1a2b3c3d5e6f7a8b7",
            "title": "Мой день рождения",
            "category": "party",
            "price": 0,
            "description": "Приглашаю вас отпраздновать мое 30-с-чем-то-летие",
            "location": {
                "city": "Москва",
                "address": "г. Москва, ул. Пушкина, дом Колотушкина"
            },
            "created_at": "2026-03-14T14:59:32+03:00",
            "created_by": "65e9c0b1a2b3c4d5e6f7a8b9",
            "started_at": "2026-04-01T12:00:00+03:00",
            "finished_at": "2026-04-01T23:00:00+03:00",
        },
    ],
    "count": 1
}
```

- `events` - список всех найденных событий
- `count` - кол-во найденных событий и должно соответствовть размеру списка *events*

**Ответ (события НЕ найдены):**

```http
HTTP/1.1 200 ОК
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Type: application/json
Content-Length: 38
```

```json
{
    "events": [],
    "count": 0
}
```

**Ответ (если параметры невалидны):**

```http
HTTP/1.1 400 Bad Request
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Length: 999
Content-Type: application/json
{"message": "invalid \"{field_name}\" field"}
```

### Мероприятие

Реализуйте новый эндпоинт `GET /events/{id}` возвращающий подробные данные о событии или мероприятии.

```http
GET /events/12e9c0b1a2b3c3d5e6f7a8b7 HTTP/1.1
Host: localhost:8000
Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d
```

**Ответ (мероприятие найдено):**

```http
HTTP/1.1 200 ОК
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Type: application/json
Content-Length: 591
```

```json
{
    "id": "12e9c0b1a2b3c3d5e6f7a8b7",
    "title": "Мой день рождения",
    "category": "party",
    "price": 0,
    "description": "Приглашаю вас отпраздновать мое 30-с-чем-то-летие",
    "location": {
        "city": "Москва",
        "address": "г. Москва, ул. Пушкина, дом Колотушкина"
    },
    "created_at": "2026-03-14T14:59:32+03:00",
    "created_by": "65e9c0b1a2b3c4d5e6f7a8b9",
    "started_at": "2026-04-01T12:00:00+03:00",
    "finished_at": "2026-04-01T23:00:00+03:00",
}
```

**Ответ (мероприятие НЕ найдено):**

```http
HTTP/1.1 404 Not Found
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Length: 0
```

### Поиск организаторов

Реализуйте новый эндпоинт `GET /users`, возвращающий список организаторов, отвечающий параметрам поиска.

```http
GET /users?name=Иван HTTP/1.1
Host: localhost:8000
Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d
```

Параметры фильтрации задаются через **GET-параметры**:

- `limit` *uint* - `(>= 0)` максимально количество в выборке (участвует в пагинации)
- `offset` *uint* - `(>= 0)` кол-во, которое нужно пропустить (участвует в пагинации)
- `name` *string* - имя организатора (поиск по вхождению, например "Иван" найдет "Иван Иванов", "Петр Иванов" и "Иван Петров")
- `id` *string* - id организатора (поиск по точному совпадению)

**Ответ (организаторы найдены):**

```http
HTTP/1.1 200 ОК
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Type: application/json
Content-Length: 203
```

```json
{
    "users": [
        {
            "id": "65e9c0b1a2b3c4d5e6f7a8b9",
            "full_name": "Иван Иванов",
            "username": "ivan_ivanov",
        }
    ],
    "count": 1
}
```

> ⚠️ Не возвращайте `password_hash`!

**Ответ (организаторы не найдены):**

```http
HTTP/1.1 200 ОК
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Type: application/json
Content-Length: 37
```

```json
{
    "users": [],
    "count": 0
}
```

**Ответ (если параметры невалидны):**

```http
HTTP/1.1 400 Bad Request
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Length: 999
Content-Type: application/json
{"message": "invalid \"{field_name}\" field"}
```

### Карточка организатора

Реализуйте новый эндпоинт `GET /users/{id}`, возвращающий подробные данные об организаторе мероприятий.

```http
GET /users/65e9c0b1a2b3c4d5e6f7a8b9 HTTP/1.1
Host: localhost:8000
Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d
```

**Ответ (организатор найден):**

```http
HTTP/1.1 200 ОК
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Type: application/json
Content-Length: 117
```

```json
{
    "id": "65e9c0b1a2b3c4d5e6f7a8b9",
    "full_name": "Иван Иванов",
    "username": "ivan_ivanov",
}
```

> ⚠️ Не возвращайте `password_hash`!

**Ответ (организатор НЕ найден):**

```http
HTTP/1.1 404 Not Found
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f97d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Length: 0
```

### Мероприятия конкретного организатора

Реализуйте новый эндпоинт `GET /users/{id}/events`, возвращающий список мероприятий по конкретному организатору

```http
GET /users/65e9c0b1a2b3c4d5e6f7a8b9/events HTTP/1.1
Host: localhost:8000
Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d
```

**Ответ (мероприятия найдены):**

```http
HTTP/1.1 200 ОК
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f97d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Type: application/json
Content-Length: 760
```

```json
{
    "events": [
        {
            "id": "12e9c0b1a2b3c3d5e6f7a8b7",
            "title": "Мой день рождения",
            "category": "party",
            "price": 0,
            "description": "Приглашаю вас отпраздновать мое 30-с-чем-то-летие",
            "location": {
                "city": "Москва",
                "address": "г. Москва, ул. Пушкина, дом Колотушкина"
            },
            "created_at": "2026-03-14T14:59:32+03:00",
            "created_by": "65e9c0b1a2b3c4d5e6f7a8b9",
            "started_at": "2026-04-01T12:00:00+03:00",
            "finished_at": "2026-04-01T23:00:00+03:00",
        }, 
    ],
    "count": 1
}
```

**Ответ (мероприятия НЕ найдены):**

```http
HTTP/1.1 200 ОК
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f97d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Type: application/json
Content-Length: 38
```

```json
{
    "events": [],
    "count": 0
}
```

### Шардирование

Реализуйте шардирование коллекции `events` по хэш-ключу `created_by` (id пользователя, который создал мероприятие).

Коллекцию `users` можно не шардировать, достаточно репликации.

### Репликация

Реализуйте репликацию с помощью Replica Set в MongoDB.

Для каждого шарда должен быть:

- 1 primary-нода
- 2 secondary-ноды

> ⚠️ На каждой ноде должен быть запущен отдельный экземпляр MongoDB,
с разными портами и разными директориями для хранения данных!

## FAQ

**Q: Зачем нужно шардирование, если у нас не так много данных?**  
**A:** Шардирование необходимо для масштабирования базы данных и обеспечения высокой доступности.
Даже если у нас сейчас всего 1000 мероприятий, в будущем их количество может значительно вырасти,
и шардирование позволит эффективно распределять нагрузку между несколькими серверами.

**Q: Нужно ли использовать Kubernetes?**  
**A:** Нет, для этой лабораторной работы достаточно запустить несколько экземпляров MongoDB в docker-контейнерах.
Главное - обеспечить правильную конфигурацию для шардирования и репликации.

**Q: Как проверить, что шардирование и репликация работают правильно?**  
**A:** Вы можете использовать MongoDB Shell или MongoDB Compass или другие инструмент
для проверки состояния шардирования и репликации.
Для шардирования проверьте, что коллекция `events` распределена между несколькими шардами.
Для репликации проверьте, что у вас есть одна primary-нода и две secondary-ноды, и что данные реплицируются между ними.

**Q: Какие значения считаются валидными для поля category?**  
**A:** Валидными значениями для поля `category` являются только те, которые указаны в условии задачи.
Любое другое значение будет считаться невалидным и должно приводить к ошибке 400 Bad Request.

**Q: Что делать, если в `PATCH /events/{id}` передано поле, которого нет в спецификации**  
**A:** Игнорировать это поле и не возвращать ошибку. Обновлять только те поля, которые указаны в спецификации.

**Q: Как обрабатывать пустую строку в поле city в `PATCH /events/{id}`?**  
**A:** Если в поле `city` передана пустая строка, это означает,
что нужно удалить `events.location.city` в MongoDB для данного мероприятия.

**Q: Как проверить, что пользователь — организатор мероприятия?**  
**A:** Для проверки, что пользователь является организатором мероприятия,
нужно сравнить `created_by` мероприятия с `id` пользователя, который делает запрос.
Если они совпадают, значит пользователь является организатором и имеет право редактировать мероприятие.
Если нет, то нужно вернуть ошибку 404 Not Found
или 401 Unauthorized, в зависимости от того, авторизован ли пользователь или нет.

**Q: Как проверить, что пользователь — организатор мероприятия?**  
**A:** Для проверки, что пользователь является организатором мероприятия,
нужно сравнить `created_by` мероприятия с `id` пользователя, который делает запрос.
Если они совпадают, значит пользователь является организатором и имеет право редактировать мероприятие.
Если нет, то нужно вернуть ошибку 404 Not Found
или 401 Unauthorized, в зависимости от того, авторизован ли пользователь или нет.
