# Набор утилит HGraber

[Ссылка](https://gitlab.com/gbh007/hgraber) на проект.

**Важно:** данный репозиторий существует для решения конкретных проблем которые возникли по мере развития проекта, он может не учитывать все потребности при работе с проектом

Примеры использования:

Экспорт данных с баз (jdb последних версий не поддерживается):

> extractor -f db.json -ft jdb -to jdb-out.json  
> extractor -f main.db -ft sqlite -to sqlite-out.json  
> extractor -f "postgres://user:pass@localhost:5432/db?sslmode=disable" -ft postgresql -to postgresql-out.json

Слияние данных с датами:

> merger -a jdb-out.json -b sqlite-out.json -to merge-out.json -diff

Исправление дат в PostgreSQL:

> datefixer -f merge-out.json -to "postgres://user:pass@localhost:5432/db?sslmode=disable"

Сканирование файловой системы для обработки актуальных дат:

> fsscan -f loads -to fsscan-out.json

Сканирование созданных файлов на соответствие данным:

> zipmatcher -f exports -src postgresql-out.json -to zip-match.json

Сканирование данных на поиск дубликатов (или на наличие данных в других):

> bookmatcher -f postgresql-out.json -src postgresql-out.json -to book-match.json -filter

## Что не поддерживает проект

1. JDB в новом формате (можно сконвертировать через PostgreSQL, но возможны потери дат)
2. Исправление дат в любых БД кроме PostgreSQL
3. Данные атрибутов и прочие расширенные данные доступны только из PostgreSQL
