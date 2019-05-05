package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Database *sql.DB

// Open - инициализация подключения к БД
func Open() error {
	if Database != nil {
		return fmt.Errorf("db: Error: Database already opened")
	}

	postgresConfig := map[string]string{
		"sslmode":  "disable",
		"host":     os.Getenv("POSTGRES_HOST"),
		"port":     os.Getenv("POSTGRES_PORT"),
		"dbname":   os.Getenv("POSTGRES_DB"),
		"user":     os.Getenv("POSTGRES_USER"),
		"password": os.Getenv("POSTGRES_PASSWORD"),
	}

	var pgConfigString string

	for key, val := range postgresConfig {
		pgConfigString += fmt.Sprintf("%s=%s ", key, val)
	}

	var err error
	Database, err = sql.Open("postgres", pgConfigString)
	if err != nil {
		return err
	}

	err = Database.Ping()
	if err != nil {
		return err
	}

	// Создание таблиц
	if query, err := ioutil.ReadFile("sql/create.sql"); err != nil {
		return err
	} else {
		if _, err := Database.Exec(string(query)); err != nil {
			return err
		}
	}
	log.Println("db: Connection: Initialized")

	return nil
}

// Close - Закрытие подключения к БД
func Close() error {
	log.Println("db: Connection: Closed")
	return Database.Close()
}
