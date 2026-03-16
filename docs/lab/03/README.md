# Лабораторная работа №3. Пользователи и события

## Подготовка

⚠️ **Обязательно ознакомьтесь с
[CONTRIBUTING.md](https://github.com/sitnikovik/ndbx-template?tab=contributing-ov-file)** -
там описан процесс работы с репозиторием, Pull Requests и GitHub Actions.

## Цель работы

Реализовать регистрацию пользователей и
создание событий, на которые в дальнейшем могут подписываться пользователи,
а для их хранения используется [MongoDB](https://www.mongodb.com/).

## Задание

### Регистрация пользователей

Реализуйте новый endpoint `POST /users` для создания пользователя, от лица которого,
можно создавать события, на которые будет подписываться все желающие, в том числе, и анонимные пользователи.

**Запрос:**

```http
POST /users HTTP/1.1
Host:localhost:8080
Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d;
Content-Type: application/json
Content-Length: 999
```

**Тело запроса:**

JSON-струтура пользователя (см. [ниже](#пользователь)) и пароль

```json
{
    "full_name": "Джон Доу",
    "username": "j0hnd0e42",
    "password": "svp4_dvp4_str0ng_passw0rd"
}
```

**Ответ (при успешном создании):**

```http
HTTP/1.1 201 Created
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Length: 0
```

**Ответ (если пользователь уже существует):**

```http
HTTP/1.1 409 Conflict
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Type: application/json
Content-Length: 999
{"message": "user already exists"}
```

**Ответ (если данные не валидны):**

```http
HTTP/1.1 400 Bad Request
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Type: application/json
Content-Length: 999
{"message": "invalid \"{field_name}\" field"}
```

### Аутентификация

Реализуйте новый endpoint `POST /auth/login` для аутентификации пользователя.

```http
POST /auth/login HTTP/1.1
Host:localhost:8080
Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/;
Content-Length: 0
```

**Тело запроса:**

```json
{
    "username": "j0hnd0e42", // Обязательный
    "password": "svp4_dvp4_str0ng_passw0rd" // Обязательный
}
```

> Необходимо проверить, что `username` и `password` присутствуют в теле запроса и не являются пустой строкой.

- `username` *string* - имя пользователя, по которому происходит аутентификация
- `password` *string* - пароль, по которому происходит аутентификация

**Ответ (при успешной аутентификации):**

```http
HTTP/1.1 204 No Content
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Length: 0
```

> ⚠️ При успешной аутентификации необходимо привязать активную сессию из куки к этому пользователю или создать новую

Добавьте в Redis в хэш-таблицу значения сессии поле `user_id` - идентификатор пользователя,
который соответствует `_id` документа в коллекции `users` в MongoDB.

**Ответ (если аутентификация не прошла):**

```http
HTTP/1.1 400 Bad Request
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
{"message": "invalid credentials"}
```

### Выход из аккаунта

Реализуйте новый endpoint `POST /auth/logout` для выхода пользователя из своего аккаунта.
Для идентификации пользователя используется сессия из куки.

> ⚠️ Критически важно удалить сессию после выхода и куки `X-Session-Id` должен быть удалён (или истечь) в ответе

```http
POST /auth/logout HTTP/1.1
Host:localhost:8080
Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d;
Content-Length: 0
```

**Ответ (при успешном выходе):**

```http
HTTP/1.1 204 No Content
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age=0
Content-Length: 0
```

> `Max-Age=0` в Set-Cookie означает, что куки должно быть удалено на стороне клиента

### Создание события

Реализуйте новый endpoint `POST /events` для создания события,
на которое могут подписываться все желающие, в том числе и не авторизованные пользователи.
Эндпоинт возвращает идентификатор созданного события,
который соответствует `_id` документа в коллекции `events` в MongoDB.

> 🔐 Доступен только для авторизованных пользователей

```http
POST /events HTTP/1.1
Host:localhost:8080
Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d;
Content-Length: 999
Content-Type: application/json
```

**Тело запроса:**

```json
{
    "title": "Мой день рождения", // Обязательный
    "address": "г. Санкт-Петербург, ул. Пушкина, дом Колотушкина", // Обязательный
    "started_at": "2026-04-01T12:00:00+03:00", // Обязательный
    "finished_at": "2026-04-01T23:00:00+03:00", // Обязательный
    "description": "Приглашаю вас отпраздновать мое 30-с-чем-то-летие",
}
```

> ⚠️ К **каждому** документу MongoDB для созданного того или иного события
> должен присваиваться идентификатор авторизованного пользователя! (см. [схему](#событие))

**Ответ (при успешной создании):**

```http
HTTP/1.1 201 OK
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Length: 999
Content-Type: application/json
{"id": "12e9c0b1a2b3c3d5e6f7a8b7"}
```

**Ответ (если пользователь не авторизован):**

```http
HTTP/1.1 401 Unauthorized
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Length: 0
```

**Ответ (если параметры невалидны):**

```http
HTTP/1.1 400 Bad Request
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Type: application/json
Content-Length: 999
{"message": "invalid \"{field_name}\" field"}
```

> ⚠️ Возвращайте ошибку, если какой-то из не валиден

**Ответ (если событие уже создано):**

```http
HTTP/1.1 409 Conflict
Set-Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d; HttpOnly; Path=/; Max-Age={APP_USER_SESSION_TTL}
Content-Type: application/json
Content-Length: 999
{"message": "event already exists"}
```

### Просмотр событий

Реализуйте новый endpoint `GET /events` для просмотра всех событий с возможностью фильтрации и пагинации.

```http
GET /events?title=my_supa_party&limit=10&offset=10 HTTP/1.1
Host:localhost:8080
Cookie: X-Session-Id=3f8a2c1d9e4b7f0a5c6d2e8b1a3f9c7d
```

Параметры фильтрации задаются через **GET-параметры**:

- `title` *string* - название события или подстрока, входящее в название события (по аналоги c `LIKE` в SQL)
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
    "events": [
        {
            "id": "12e9c0b1a2b3c3d5e6f7a8b7",
            "title": "Мой день рождения",
            "description": "Приглашаю вас отпраздновать мое 30-с-чем-то-летие",
            "location": {
                "address": "г. Санкт-Петербург, ул. Пушкина, дом Колотушкина"
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
Content-Length: 999
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
{"message": "invalid \"{field_name}\" parameter"}
```

### Пользователь

**Коллекция:** `users`

**Схема:**

```json
{
    "full_name": "Джон Доу",
    "username": "j0hnd0e42",
    "password_hash": "$2a$10$Xr0DDNUTfpbLihAp0ZbGPei1oFP8g5FNypIvaXdW7W.KWJaobPA5q"
}
```

> ⚠️ Храните именно **хэш пароля** по алгоритму **bcrypt** вместо самого пароля, по которому авторизуется пользователь!

**Индексы:**

*Ожидается, как минимум, один индекс:*

- `username` *unique* - уникальный индекс по юзернеймам пользователей, по которому происходит аутентификация

### Событие

**Коллекция:** `events`

**Схема:**

```json
{
    "title": "Мой день рождения",
    "description": "Приглашаю вас отпраздновать мое 30-с-чем-то-летие",
    "location": {
        "address": "г. Санкт-Петербург, ул. Пушкина, дом Колотушкина"
    },
    "created_at": "2026-03-14T14:59:32+03:00",
    "created_by": "65e9c0b1a2b3c4d5e6f7a8b9", // hex из id пользователя, который создал событие (авторизованный)
    "started_at": "2026-04-01T12:00:00+03:00",
    "finished_at": "2026-04-01T23:00:00+03:00",
}
```

> ⚠️ Все даты храним только в строке формата RFC3339 с часовым поясом (может быть любым)

**Индексы:**

*Ожидается, как минимум, такой набор индексов:*

- `title` *unique* - уникальный индекс по названиям событий
- `title, created_by` - составной индекс по названию и автору
- `created_by` - индекс по автору

### Конфигурация

Добавьте следующие переменные в `.env.local`:

```sh
# Название базы данных в MongoDB
MONGODB_DATABSE="eventhub" # или любое другое на ваше усмотрение
# Имя пользователя для аутентификации в MongoDB
MONGODB_USER=your_mongodb_username
# Пароль для аутентификации в MongoDB
MONGODB_PASSWORD=your_mongodb_password
# Хост MongoDB сервера
MONGODB_HOST=your_mongodb_container_name
# Порт MongoDB сервера
MONGODB_PORT=your_mongodb_container_port
```

> ⚠️ Для тестов рекомендуется оставить `APP_USER_SESSION_TTL` небольшим, чтобы ускорить проверку истечения сессии,
как это требовалось в [предыдущей лабораторной работе](/docs/lab/02/)

## FAQ

**Q: Когда обновляется сессия?**  
A: На каждом POST-запросе. GET запросы только возвращают существующую сессию.

**Q: Что произойдет когда истечет сессия?**  
A: Авторизация должна "слететь" и пользователь не сможет создать события, пока не авторизуется заново.
При этом, он может просмотривать события через `GET /events`.

**Q: Откуда берется поле id в `GET /events?`**  
A: Поле `id` в ответе `GET /events` соответствует `_id` документа в коллекции `events` в MongoDB.

**Q: MongoDB позволяет хранить "что угодно". Можно ли поменять схему?**  
A: Можете дополнять, но менять формат данных или названия полей нельзя, так как это сломает АТ.

**Q: Можно делать больше индексов чем требуется?**  
A: Еще успеете 😉

**Q: Нужно ли реализовывать удаление и изменение событий?**  
A: Не обязательно. Это будет в следующей лабораторной работе.

**Q: Могу ли я использовать camelCase в названиях полей документов и в запросах?**  
A: Нет. Только snake_case, иначе ваша работа не пройдет АТ.

**Q: Что на счет `POST /session` из [предыдущей лабораторной](/docs/lab/02/)?**  
A: Ничего. В рамках этого задания этот эндпоинт не используется.

**Q: Кто угодно может регистрировать пользователей?**  
A: В нашем проекте - да, но для регистрации необходима сессия.
Мы опускаем момент с верификацией, двухфакторной аутентификацией и прочими принципами защиты от злоумышленников,
так как это академический проект и не претендует на звание самого безопасного проекта в мире 😉

**Q: Можно не использовать логин/пароль для MongoDB?**  
A: Можно. Но, лучше не надо.
