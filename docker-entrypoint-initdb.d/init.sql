-- Создание базы данных
CREATE DATABASE my_app_db;

-- Использование базы данных
\c my_app_db;

-- Создание таблицы sessions
CREATE TABLE IF NOT EXISTS sessions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    active BOOLEAN DEFAULT true
);

-- Пример вставки данных
INSERT INTO sessions (user_id, active) VALUES (1, true);