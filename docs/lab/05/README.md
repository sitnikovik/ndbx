# Лабораторная работа №5

## Подготовка

⚠️ **Обязательно ознакомьтесь с
[CONTRIBUTING.md](https://github.com/sitnikovik/ndbx-template?tab=contributing-ov-file)** -
там описан процесс работы с репозиторием, Pull Requests и GitHub Actions.

## Цель работы

Реализовать функционал реакций на мероприятия и их организаторов,
используя [Apache Cassandra](https://cassandra.apache.org/)
в качестве хранилище для них
и [Redis](https://redis.io/) для их кэширования.

## Эндпоинты

### Лайк на мероприятие

Реализуйте новый эндпоинт `POST /events/{event_id}/like`,
который ставит лайк пользователя на понравившееся мероприятие.

> 🔐 Доступно  **только авторизованным** пользователям

**Запрос:**

```http
POST /events/{event_id}/like HTTP/1.1
Host: localhost:8000
Content-Type: application/json
Content-Length: 0
Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d;
```

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

### дизлайк на мероприятие

Реализуйте новый эндпоинт `POST /events/{event_id}/dislike`,
который ставит дизлайк пользователя на конкретное мероприятие.

> 🔐 Доступно  **только авторизованным** пользователям

**Запрос:**

```http
POST /events/{event_id}/dislike HTTP/1.1
Host: localhost:8000
Content-Type: application/json
Content-Length: 0
Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d;
```

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
Set-Cookie: X-Session-Id=; HttpOnly; Path=/; Max-Age=0
Content-Length: 0
```

### Реакции в мероприятиях

Доработайте существующие эндпоинты `GET /events`, `GET /events/{event_id}`, `GET /users/{user_id}/events` так,
чтобы при переданном параметре `include=reactions`,
в каждом мероприятии возвращались счетчики всех его лайков (фильтры не применяются к ним)

```http
GET /events?include=reactions&limit=10 HTTP/1.1
Host:localhost:8080
Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d
```

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
            "reactions": {
                "likes": 24,
                "dislikes": 3,
            }
        },
    ],
    "count": 1
}
```

> ⚠️ Все эндпоинты по мероприятиям должны возврщать объект `reactions` (при `?include=reactions` в запросе)
даже если у мероприятия нет ни одной реакции/лайка как в Cassandra, так и Redis

#### Подсчет реакций

Для вычисления кол-ва реакций для мероприятия нужно отобрать все-все встречи по их названиям,
просуммировать реакции по ним и вернуть счетчики в endpoint'ах.

Например, есть три спектакля "Алиса в стране чудес" на три разных времени.
В MongoDB это создано как три разных сущности, но для пользователя они считается за одно мероприятие,
которое проводится несколько раз в отведенное время.
И пользователь должен видеть счетчик реакция как на одно мероприятие, а не на три разных.

## Спецификация

### Лайк мероприятия

Лайки хранятся в Cassandra и Redis. Cassandra как основное хранилище, а Redis для кэширования чтения.
Реализуйте подход [Cache-Aside](https://habr.com/ru/articles/991332/),
при котором Redis будет "заполнятся" приложениям при запросе лайков из Cassandra.

**TTL для лайков** в кэше задается через env-конфиг:

```sh
# Время жизни лайков в кэше в секундах.
APP_LIKE_TTL=60
```

> 💡 Так обеспечивается высокая доступность и быстрое чтение лайков

#### Cassandra

`event_reactions` - название таблицы реакций на мероприятия

**Схема**:

- `event_id` *text* - идентификатор мероприятия в MongoDB
- `like` *tinyint* - значение записи
  - `1` - лайк
  - `-1` - дизлайк
- `created_by` *string* - идентификатор пользователя в MongoDB, который оставил реакцию
- `created_at` *timestamp* - дата и время создания реакция (в UTC)

> ⚡️ Не ищите лайк, чтобы проверить, что он есть - удаляйте и пишите новую запись

#### Redis

**Ключ:**

```plaintext
events:{event_title_md5_hash}:reactions
```

`{event_title_md5_hash}` - хэш по алгоритму md5 для названия мероприятия.

Например, для "Алиса в стране чудес" ключ будет

```plaintext
events:ba80405c3ebccb9cb99791b47c2487f9:reactions
```

**Значение:**

```json
{
    "likes": 24,
    "dislikes": 3
}
```

### Конфигурация

Добавьте следующие переменные в `.env.local`:

```sh
# Список хостов Cassandra в виде строки, разделенной запятыми (пока что один хост).
CASSANDRA_HOSTS=cassandra-test
# Порт Cassandra
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

**Q: Нужно ли заполнять кэш лайками, если нет соответствующей записи в Cassandra?**  
A: Нет. В кэш пишем только данные только по существующим данным.
Но так или иначе, в эндпоинтах [возвращаем](#реакции-в-мероприятиях) `reactions` c нулями.

**Q: Что если несколько раз отправить `POST /events/{event_id}/like?**  
A: Если лайк уже был - произойдет удаление и создание новой записи.

**Q: Что если встреч с одинаковым названием несколько, а реакции "привязываются" к мероприятиям по их идентификаторам?
A: В эндпоинтах возвращаются все-все реакции по мероприятиям, но в Cassandra они записываются как к разным сущностям.
Вы должны отобрать все встречи по их названиям и посчитать все реакции от всех пользователей.

**Q: Нужно ли подготовить Cassandra перед использованием?**
A: Да, необходимо создать ключевое пространство и таблицу `event_reactions` согласно спецификации выше при запуске приложения через `make run`.
