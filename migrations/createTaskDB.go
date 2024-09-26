package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	// "gorm.io/driver/postgres"
)

// Функция для выполнения SQL-скрипта из файла
func ExecuteSQLFile(dsn string, filePath string) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}
	defer db.Close()

	// Чтение содержимого SQL-файла
	sqlBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("failed to read SQL file: %v", err)
	}

	// Выполнение SQL-команд
	_, err = db.Exec(string(sqlBytes))
	if err != nil {
		log.Fatalf("failed to execute SQL script: %v", err)
	}

	fmt.Println("Tables created successfully.")
}
