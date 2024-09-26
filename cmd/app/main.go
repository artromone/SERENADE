package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"serenade/migrations"

	_ "github.com/lib/pq"
	// "gorm.io/driver/postgres"
)

func main() {
	// Указываем DSN для подключения к PostgreSQL без указания базы данных
	dsn := os.Getenv("DATABASE_URL")

	// Подключаемся к PostgreSQL
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer db.Close()

	// Создаем базу данных, если она не существует
	_, err = db.Exec("CREATE DATABASE task_database")
	if err != nil {
		if err.Error() != "pq: database \"task_database\" already exists" {
			log.Fatalf("failed to create database: %v", err)
		}
		fmt.Println("Database already exists.")
	} else {
		fmt.Println("Database created successfully.")
	}

	// Теперь подключаемся к созданной базе данных
	dsn += " dbname=task_database"

	// Выполнение SQL-скрипта для создания таблиц
	migrations.ExecuteSQLFile(dsn, "migrations/create_db_and_tables.sql")
}
