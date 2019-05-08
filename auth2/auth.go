package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"2019_1_qwerty/auth2/auth"

	_ "github.com/lib/pq"
	"golang.org/x/net/context"
)

var Database *sql.DB

func Open() error {
	if Database != nil {
		log.Fatalln("db: Error: Database already opened")
		return nil
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
	log.Println("db: Connection: Initialized")
	return nil
}

const sessKeyLen = 10

type SessionManager struct {
	sessions map[string]*auth.User
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions: map[string]*auth.User{},
	}
}

const sqlSelectUserPasswordByNickname = `
SELECT password
FROM users
WHERE nickname = $1
`

func (sm *SessionManager) LoginUser(ctx context.Context, in *auth.User) (*auth.Status, error) {
	var dbPassw string
	row := Database.QueryRow(sqlSelectUserPasswordByNickname, in.Nickname)
	if err := row.Scan(&dbPassw); err != nil {
		return &auth.Status{Ok: false}, nil
	}

	if (in.Password) != dbPassw {
		return &auth.Status{Ok: false}, nil
	}
	return &auth.Status{Ok: true}, nil
}
