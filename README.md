# NoSQL Database Exploration

[![Quality](https://github.com/sitnikovik/ndbx/actions/workflows/quality.yml/badge.svg)](https://github.com/sitnikovik/ndbx/actions/workflows/quality.yml)
![Lab Progress](https://img.shields.io/badge/Lab-1%20of%207-blue)

Проект содержит все необходимые материалы для курса по NoSQL базам данных и
пайплайн в GitHub Actions для проверки лабораторных работ.

## Лабораторный проект

**EventHub** — backend‑сервис платформы мероприятий, предназначенный для изучения различных подходов
к хранению и обработке данных с использованием NoSQL баз данных.
Проект выполняется поэтапно в рамках лабораторных работ
и представляет собой единый сервис, который развивается в течение всего курса.

**Цель проекта** — на практике понять, как работают NoSQL базы данных,
их сильные и слабые стороны, и сравнить различные подходы к хранению данных.

Вы научитесь работать с 4-мя NoSQL базами данных:

- [Redis](https://redis.io/)
- [MongoDB](https://www.mongodb.com/)
- [Cassandra](https://cassandra.apache.org/)
- [Neo4j](https://neo4j.com/)

Всего будет **7 лабораторных работ**:

1. [Старт: Healthcheck](docs/lab/01/)
2. Redis: Сессии пользователей
3. MongoDB: Документные данные
4. MongoDB: Шардирование и репликация
5. Cassandra: Агрегация и просмотры
6. Cassandra: Репликация событий
7. Neo4j: Связи и рекомендации

> Задания публикуются постепенно по мере прохождения курса

## Требования к реализации

- Используйте предоставленный шаблон репозитория
[ndbx-template](https://github.com/sitnikovik/ndbx-template)
- Соблюдайте требования к структуре проекта, описанные в
[CONTRIBUTING.md](https://github.com/sitnikovik/ndbx-template?tab=contributing-ov-file)
- Можно использовать **любой язык программирования**
- Фреймворки и библиотеки - на ваше усмотрение, но не забывайте о производительности
- Все лабораторные работы должны полностью проходить
автоматическую проверку в GitHub Actions
- Каждая последующая лабораторная работа должна
на основе предыдущей
- Очередная лабораторная работа не должна ломать
функциональность предыдущих лабораторных работ
(проверяется в процессе автоматической проверки)

## Система оценивания

За выполнение лабораторных работ вы получаете баллы,
из которых складывается итоговая оценка за семестр.

У каждой лабораторной работы есть своя максимальная сумма баллов и **срок сдачи**.
Если вы сдаёте работу после срока,
то из полученных баллов вычитается "штраф за просрочку".

Сроки и баллы указаны в таблице успеваемости (уточняйте у преподавателя).

## Рекомендации

- **Документируйте** исходный код, соблюдая стандарты, например:
  - [Javadoc](https://docs.oracle.com/en/java/javase/17/docs/specs/javadoc/doc-comment-spec.html)
  - [PEP8](https://peps.python.org/pep-0008/)
  - [Go doc](https://go.dev/doc/comment)
  - [PHPDoc](https://github.com/php-fig/fig-standards/blob/master/proposed/phpdoc.md)
  - [JSDoc](https://jsdoc.app/)
- Собюдайте [конвенцию коммитов](https://www.conventionalcommits.org/ru/v1.0.0-beta.2/)
- Каждый коммит - рабочая и протестированная версия проекта. Это облегчит поиск последней рабочей версии при поиске ошибки
- Пишите документацию на [Markdown](https://www.markdownguide.org/) в [Obsidian](https://obsidian.md/) или в вашей IDE.
- Используйте [Makefile](https://makefiletutorial.com/) для удобного взаимодействия с проектом
- Используйте [PlantUML](https://plantuml.com/ru/) для создания блок-схемы и диаграмм.
Есть возможность предпросмотра в IDE
(в VSCode через [расширение PlantUML](https://marketplace.visualstudio.com/items?itemName=jebbs.plantuml))
и утилиты graphviz и Java, установленных на вашем компьютере ([подробнее](https://plantuml.com/ru-dark/faq-install))
- Схемы БД реализуйте через [DBML](https://dbml.dbdiagram.io/home/)
и [dbdiagram.io](https://dbdiagram.io/home).
В VSCode есть [расширение](https://marketplace.visualstudio.com/items?itemName=dbdiagram.dbdiagram-vscode) для предпросмотра
- Ведите коллекцию [Postman](https://www.postman.com/)
или [Insomnia](https://insomnia.rest/) c примерами запросов и ответов для **каждого эндпоинта**.
Можно развернуть [Swagger](https://swagger.io/), но не обязательно.
- Используйте публичные **датасеты** с тестовыми данными,
например, из [Kaggle](https://www.kaggle.com/) или генерируйте свои на Python, Go или любом другом языке программирования.
- Если нужен текст-рыба - используйте
  - [генератор текста "рыба"](https://fish-text.ru/) если нужен текст на русском
  - или [Lorem Ipsum](https://getlorem.com/) если нужен текст на английском
