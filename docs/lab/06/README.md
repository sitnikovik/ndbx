# Лабораторная работа №5

## Подготовка

⚠️ **Обязательно ознакомьтесь с
[CONTRIBUTING.md](https://github.com/sitnikovik/ndbx-template?tab=contributing-ov-file)** -
там описан процесс работы с репозиторием, Pull Requests и GitHub Actions.

## Цель работы

Масштабировать кластер [Apache Cassandra](https://cassandra.apache.org/)
и реализовать функционал отзывов на мероприятия.

## Эндпоинты

### Оставить отзыв на мероприятие

Реализуйте новый эндпоинт `POST /events/{event_id}/reviews`,
с помощью которого пользователи могут оставлять отзывы и комментарии на мероприятия.

> 🔐 Доступно **только авторизованным** пользователям

**Запрос**:

```http
POST /events/{event_id}/reviews HTTP/1.1
Host:localhost:8080
Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d;
Content-Type: application/json
Content-Length: 999
```

**Тело запроса**:

```json
{
    "comment": "Великолепный спектакль! Идите, даже не думайте!",
    "rating": 5
}
```

- `comment` *string* - комментарий к отзыву (любые символы, но максимум 300)
- `rating` *int* - оценка от `1` до `5` (только **целые** числа)

**Ответ (OK):**

```http
HTTP/1.1 204 No Content
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Length: 0
```

**Ответ (уже добавлял отзыв):**

```http
HTTP/1.1 409 No Content
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Length: 0
{"message": "Already exists"}
```

**Ответ (мероприятие не найдено или нет доступа к нему):**

```http
HTTP/1.1 404 Not Found
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Type: application/json
Content-Length: 29
{"message": "Event not found"}
```

**Ответ (если параметры невалидны):**

```http
HTTP/1.1 400 Bad Request
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Length: 999
Content-Type: application/json
{"message": "invalid \"{field_name}\" field"}
```

**Ответ (если пользователь не авторизован):**

```http
HTTP/1.1 401 Unauthorized
Content-Length: 0
```

> ⚠️ Одно мероприятие - один пользователь — один отзыв

### Просмотр отзывов на мероприятие

Создайте новый эндпоинт `GET /events/{event_id}/reviews`,
возвращающий отзывы для конкретного мероприятия c пагинацией.

```http
GET /events/{event_id}/reviews?limit=10&offset=10 HTTP/1.1
Host:localhost:8080
```

Параметры задаются через **GET-параметры**:

- `limit` *uint* - `(>= 0)` максимально количество событий в выборке (участвует в пагинации)
- `offset` *uint* - `(>= 0)` кол-во событий, которое нужно пропустить (участвует в пагинации)
**Ответ (события найдены):**

```http
HTTP/1.1 200 ОК
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Type: application/json
Content-Length: 999
```

```json
{
    "reviews": [
        {
            "id": "56e2c0b3a2b4c1a5e6f7f8b3",
            "event_id": "12e9c0b1a2b3c3d5e6f7a8b7",
            "comment": "Великолепный спектакль! Идите, даже не думайте!",
            "created_at": "2026-03-14T14:59:32+03:00",
            "created_by": "65e9c0b1a2b3c4d5e6f7a8b9",
            "rating": 5,
            "updated_at": "2026-03-14T14:59:32+03:00"
        },
    ],
    "count": 1
}
```

- `reviews` - список всех найденных отзывов
- `count` - кол-во найденных отзывов и должно соответствовть размеру списка *reviews*

**Ответ (отзывы НЕ найдены):**

```http
HTTP/1.1 200 ОК
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Type: application/json
Content-Length: 38
```

```json
{
    "reviews": [],
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

### Изменить отзыв на мероприятие

Реализуйте новый эндпоинт `PATCH /events/{event_id}/reviews/{review_id}`,
с помощью которого пользователи могут редактировать отзывы и комментарии на мероприятия.

> 🔐 Доступно **только авторизованным** пользователям и только владельцам самих отзывам

**Запрос**:

```http
PATCH /events/{event_id}/reviews/{review_id} HTTP/1.1
Host:localhost:8080
Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d;
Content-Type: application/json
Content-Length: 999
```

**Тело запроса**:

```json
{
    "rating": 3,
    "comment": "На самом деле, так себе спектакль..."
}
```

- `rating` *int* - оценка от `1` до `5` (только **целые** числа)
- `comment` *string* - комментарий к отзыву (любые символы, но максимум 300)

> 💡 Эндпоинт изменяет в отзыве только те поля, которые переданы
и **всегда** обновляет поле `updated_at` на текущее время при успешном обновлении отзыва

**Ответ (OK):**

```http
HTTP/1.1 204 No Content
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Length: 0
```

**Ответ (мероприятие не найдено или нет доступа к нему):**

```http
HTTP/1.1 404 Not Found
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Type: application/json
Content-Length: 29
{"message": "Event not found"}
```

**Ответ (если пользователь не авторизован):**

```http
HTTP/1.1 401 Unauthorized
Content-Length: 0
```

### Отзывы в мероприятиях

Доработайте существующие эндпоинты `GET /events`, `GET /events/{event_id}`, `GET /users/{user_id}/events` так,
чтобы при переданном параметре `include=reviews`,
в каждом мероприятии возвращалась усредненная информация об отзывах (фильтры не применяются к ним)

```http
GET /events?include=reviews&limit=10 HTTP/1.1
Host:localhost:8080
Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d
```

> 💡 В `include` можно передать несколько значений через **запятую**, например, `?include=reactions,reviews`

**Ответ (события найдены):**

```http
HTTP/1.1 200 ОК
Content-Type: application/json
Content-Length: 999
```

```json
{
    "events": [
        {
            "id": "12e9c0b1a2b3c3d5e6f7a8b7",
            // ...
            "reviews": {
                "count": 123,
                "rating": 4.8
            },
        },
    ],
    "count": 1
}
```

> ⚠️ Все эндпоинты по мероприятиям должны возвращать объект `reviews` (при `?include=reviews` в запросе)
даже если у мероприятия нет ни одной отзыва как в Cassandra, так и Redis

## Спецификация

### Отзыв на мероприятие

#### Cassandra

`event_reviews` - название таблицы отзывов на мероприятия

**Схема**:

- `id` *uuid* - идентификатор отзыва
- `event_id` *text* - идентификатор мероприятия в MongoDB
- `rating` *tinyint* - оценка от `1` до `5` (только **целые** числа)
- `comment` *text* - комментарий (любые символы)
- `created_by` *string* - идентификатор пользователя в MongoDB, который оставил реакцию
- `created_at` *timestamp* - дата и время создания отзыва (в UTC)
- `updated_at` *timestamp* - дата и время обновления отзыва (в UTC)

> ⚡️ Рекомендуется добавить `created_at` по **убыванию** в ключ кластеризации

#### Redis

**Ключ:**

```plaintext
events:{event_title_md5_hash}:reviews
```

`{event_title_md5_hash}` - хэш по алгоритму md5 для названия мероприятия.

Например, для "Алиса в стране чудес" ключ будет

```plaintext
events:ba80405c3ebccb9cb99791b47c2487f9:reviews
```

**Значение:**

```json
{
    "count": 123,
    "rating": 4.8
}
```

- `count` *int* - общее кол-во всех отзывов на мероприятие (по названию)
- `rating` *float* - средний рейтинг всех отзывов на мероприятие (по названию)
с округлением десятой части в большую сторону

### Конфигурация

Добавьте еще 2 хоста Cassandra в конфигурацию приложения и измените уровень консистентности

```sh
# Список хостов Cassandra в виде строки, разделенной запятыми.
CASSANDRA_HOSTS=cassandra-1,cassandra-2,cassandra-3,
# Порт основного узла
CASSANDRA_PORT=9042
# Имя пользователя для подключения к Cassandra
CASSANDRA_USERNAME=
# Пароль для подключения к Cassandra
CASSANDRA_PASSWORD=
# Имя ключевого пространства Cassandra
CASSANDRA_KEYSPACE="testkeyspace"
# Уровень согласованности Cassandra
CASSANDRA_CONSISTENCY="QUORUM"
```

## FAQ

**Q: Можно ли порты Cassandra передавать вместе с хостами?**  
A: Можно, но в `CASSANDRA_PORT` описывает порт основного узла,
к которому обращаются АТ для выполнения запросов.

**Q: Что если отправить `POST /events/{event_id}/reviews` а несколько раз подряд?**  
A: Один пользователь — один отзыв.
Если отзыв уже был создан — возвращаем 409 ошибку, сообщающая, что отзыв уже был создан.

**Q: Редактировать отзыв через `PATCH` можно сколь угодно раз?**  
A: У нас — да. Мы опускаем момент с модерацией и прочей бизнес-логикой.

**Q: Можно ли сортировать отзывы в `GET /events/{event_id}/reviews`?**  
A: Нет. Отзывы лежат в Cassandra и сортировка "на лету" сильно замедлит время ответа,
поэтому такой функционал должен быть реализован клиентом.
Это абсолютно нормально. Тем более, что есть возможность ограничения выборки
через GET-параметры `limit` и `offset`.

**Q: GET-параметр `include` может включать в себя несколько значений?**  
A: Да. В запросах получения мероприятий можно запросить информацию по лайкам и отзывам
и значений передаются через **запятую**. Например, `?include=reactions,reviews`
