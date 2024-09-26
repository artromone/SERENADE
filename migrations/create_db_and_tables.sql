-- create_db_and_tables.sql

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'task_database') THEN
        EXECUTE 'CREATE DATABASE task_database';
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS lists (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    list_id INT REFERENCES lists(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(50),
    start_time TIMESTAMP,
    end_time TIMESTAMP,
    priority VARCHAR(50)
);

