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

## Сборка и запуск Docker-образа
- Сборка проекта
```bash
 docker build -t go-webrtc-backend .
```

## Выполнение строки запроса
```bash
curl -v -X POST http://10.64.140.43      -d "id=102"      -d "user=ivan"
```

## Просмотр содержимого таблицы
```bash
sudo microk8s kubectl exec -it -n go-webrtc deployment/postgres -- psql -U go_user -d my_app_db -c "SELECT * FROM sessions;"
```


