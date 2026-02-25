Задача:
Создать программу на go которая создаёт сессии в Postgresql. Создать для этого GRPC-сервер. Сделать взаимодействие  сложных структур и интерфейсов 

## Реализация
В данном случае Postgresql будет установлен на компьютер.
Для базы данных создана простейшая таблица:
```SQL
postgres=# create database my_app_db
postgres-# ;
CREATE DATABASE

my_app_db=# create table sessions (
id SERIAL PRIMARY KEY,
user_id integer,
active boolean default true)
;
```
Получилась такая таблица:
my_app_db=# \d sessions
                             Table "public.sessions"
 Column  |  Type   | Collation | Nullable |               Default                
---------+---------+-----------+----------+--------------------------------------
 id      | integer |           | not null | nextval('sessions_id_seq'::regclass)
 user_id | integer |           |          | 
 active  | boolean |           |          | true
Indexes:
    "sessions_pkey" PRIMARY KEY, btree (id)

