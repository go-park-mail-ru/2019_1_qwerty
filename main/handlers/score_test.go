package handlers

import (
	"2019_1_qwerty/database"
	"2019_1_qwerty/helpers"
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

func init() {
	var err error
	_ = godotenv.Load()
	helpers.RedisConnect, err = redis.DialURL("redis://redis@localhost:6379/0")
	if err != nil {
		log.Println(err)
	}
	log.Println("Redis: Connection: Initialized")
	database.Database, err = sql.Open("sqlite3", ":memory:")

	if err != nil {
		log.Fatal(err)
	}
	if _, err := database.Database.Exec(createUserTable); err != nil {
		log.Println(err)
	}
	log.Println("db: Connection: Initialized")
}

func TestScoreCreate(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/score?points=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateScore)

	handler.ServeHTTP(rr, req)

	expectedStatus := http.StatusCreated
	if status := rr.Code; status != expectedStatus {
		t.Errorf("Неожиданный код ответа: получено %v, ожидалось %v",
			status, expectedStatus)
	}
}

func TestGetScore(t *testing.T) {
	req, err := http.NewRequest("Get", "/api/score", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetNextAfter)

	handler.ServeHTTP(rr, req)

	expectedStatus := http.StatusOK
	if status := rr.Code; status != expectedStatus {
		t.Errorf("Неожиданный код ответа: получено %v, ожидалось %v",
			status, expectedStatus)
	}
}
