package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

// // Database structure
// type Database struct {
// 	Pool   *pgx.ConnPool
// 	Status models.Status
// }

// // Instance of database
var Database *sql.DB

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgresConfig := map[string]string{
		"sslmode":  "disable",
		"host":     os.Getenv("POSTGRES_HOST"),
		"port":     os.Getenv("POSTGRES_PORT"),
		"dbname":   os.Getenv("POSTGRES_DB"),
		"user":     os.Getenv("POSTGRES_USER"),
		"password": os.Getenv("POSTGRES_PASSWORD"),
	}

	var DB_CONNECT_STRING string
	for key, val := range postgresConfig {
		DB_CONNECT_STRING += fmt.Sprintf("%s=%s ", key, val)
	}

	db, err := sql.Open("postgres", DB_CONNECT_STRING)

	if err != nil {
		log.Fatal("db: Error: The data source arguments are not valid")
	}

	Database = db

	err = Database.Ping()

	if err != nil {
		log.Fatal("db: Error: Could not establish a connection with the database")
	}

	log.Println("db: Connection: OK!")
}
