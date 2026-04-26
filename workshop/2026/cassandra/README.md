# Воркшоп по Cassansdra

## С чего начать

**1. Запуск контейнеров:**

Сначала запустить первую ноду

```sh
docker compose up -d cassandra-dc1-node1
```

И убедитесь, что она полностью готова (должна быть **UN**)

```sh
docker compose exec cassandra-dc1-node1 nodetool status
```

> ⚠️ Хэсчека в Docker может быть недостаточно

И запустите остальные

```sh
docker compose up -d cassandra-dc1-node2 cassandra-dc1-node3
```

**2. Настраиваем Сassandra:**

```sh
docker compose exec -T cassandra-dc1-node1 cqlsh < ./scripts/cql/init.cql
```

**3. Генератор данных:**

Для генерации данных в таблицу `iot.sensor_data` есть CLI-команда, написанная на Go (код [здесь](./sensorgen/))

Рекомедуется скачать готовый бинарник `sensorgen` [из релиза](https://github.com/sitnikovik/ndbx/releases/tag/v1.5.0), чтобы не собирать его самим, если у вас нет Go на компьютере.

**Запуск:**

```sh
./sensorgen -n=5 -hosts=localhost -port=9042 -keyspace=iot -consistency=QUORUM
```

Посмотреть доступные флаги:

```sh
./sensorgen -h
```

## Попробуйте сами

### Zombie Data

**1. Остановите ноду:**

```sh
docker compose stop cassandra-dc1-node3
```

**2. На живой ноде вставьте:**

```sql
CONSISTENCY ONE;

INSERT INTO iot.sensor_data (sensor_id, created_at, value)
VALUES ('sensor_zombie', toTimestamp(now()), 999);
```

**3. Запуститие ноду:**

```sh
docker compose start cassandra-dc1-node3
```

**4. Вставьте старую версию:**

```sql
CONSISTENCY QUORUM;
UPDATE iot.sensor_data 
USING TIMESTAMP 1000 
SET value = 1000 
WHERE sensor_id = 'sensor_zombie' AND created_at = '2026-04-25 10:00:00+0000';
```

**5. Читаем с ONE и видим разные значения:**

```sql
CONSISTENCY ONE;
SELECT * FROM iot.sensor_data WHERE sensor_id = 'sensor_zombie';
```

**6. Чиним:**

```sql
nodetool repair iot
```

### LWT — Light-weight Transactions

Lightweight Transaction (LWT) — это способ выполнить условное обновление в Cassandra с помощью IF.

Он использует Paxos-протокол для обеспечения линеаризуемости и позволяет избежать гонок (race conditions).

**Запустите:**

> ⚠️ Скачайте [lwt из релиза](https://github.com/sitnikovik/ndbx/releases) себе на компьютер

```sql
make lwt
```

Демонстрирует использование LWT для безопасного обновления статуса сенсора
**Зачем LWT?**

- Два сервиса могут одновременно решить изменить статуса sensor_1
- Без условия IF: оба обновления пройдут → статус будет меняться дважды
- С IF last_seen_unix < ?: только один запрос применится
- Это гарантирует, что обновление происходит не чаще, чем раз в N секунд

> ⚠️ Обратите внимание: демонстрация LWT в cqlsh покажет,
> как работает условное обновление, но не покажет реальную гонку,
> потому что запросы выполняются последовательно.
>
> Чтобы увидеть настоящую пользу от LWT, **нужно создать конкурентную среду** — например, запустить два параллельных клиента.

**1. Создайте таблицу:**

```sql
CREATE TABLE IF NOT EXISTS iot.sensor_status (
    sensor_id TEXT PRIMARY KEY,
    last_seen_unix BIGINT,
    status TEXT
);
```

**2.  Вставьте начальные данные:**

```sql
-- 1777111200 = Sat Apr 25 2026 10:00:00 GMT+0000
INSERT INTO iot.sensor_status (sensor_id, last_seen_unix, status)
VALUES ('sensor_1', 1777111200, 'online');

-- 1777111500 = Sat Apr 25 2026 10:05:00 GMT+0000
INSERT INTO iot.sensor_status (sensor_id, last_seen_unix, status)
VALUES ('sensor_2', 1777111500, 'online');
```

**3. Проверьте данные:**

```sql
SELECT * FROM iot.sensor_status;
```

**4. Обновите статус ТОЛЬКО если прошло больше 5 минут:**

```sql
UPDATE iot.sensor_status 
SET last_seen_unix = 1777111500,
    status = 'offline' 
WHERE sensor_id = 'sensor_1' 
IF last_seen_unix < 1777111500;
```

Выведет:

- `applied = true`: условие выполнилось и обновление произошло
- `applied = false`: условие не выполнилось и обновление не произошло
