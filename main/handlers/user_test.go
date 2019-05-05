package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"2019_1_qwerty/database"
	"2019_1_qwerty/helpers"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

const createUserTable = `
CREATE TABLE IF NOT EXISTS users (
    nickname        CITEXT PRIMARY KEY,
    email           CITEXT UNIQUE,
    "password"  text,
    avatar          TEXT DEFAULT 'default0.jpg'
);`

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
func TestUser(t *testing.T) {
	data, _ := json.Marshal(map[string]string{"nickname": "test", "email": "Test@test.ru", "password": "Test"})
	buf := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", "/user/create", buf)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)

	handler.ServeHTTP(rr, req)

	expectedStatus := http.StatusCreated
	if status := rr.Code; status != expectedStatus {
		t.Errorf("Неожиданный код ответа: получено %v, ожидалось %v",
			status, expectedStatus)
	}

	if rr.HeaderMap["Set-Cookies"] != nil {
		t.Errorf("Cookie не были установлены после регистрации!")
	}

	req, err = http.NewRequest("GET", "/user/logout", buf)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(LogoutUser)

	handler.ServeHTTP(rr, req)

	expectedStatus = http.StatusOK
	if status := rr.Code; status != expectedStatus {
		t.Errorf("Неожиданный код ответа: получено %v, ожидалось %v",
			status, expectedStatus)
	}

	data, _ = json.Marshal(map[string]string{"nickname": "test", "password": "Test"})
	buf = bytes.NewBuffer(data)
	req, err = http.NewRequest("POST", "/user/login", buf)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(LoginUser)
	handler.ServeHTTP(rr, req)

	expectedStatus = http.StatusOK
	if status := rr.Code; status != expectedStatus {
		t.Errorf("Неожиданный код ответа: получено %v, ожидалось %v",
			status, expectedStatus)
	}

	if rr.HeaderMap["Set-Cookies"] != nil {
		t.Errorf("Cookie не были установлены после регистрации!")
	}

	data, _ = json.Marshal(map[string]string{"nickname": "test", "password": "Test"})
	buf = bytes.NewBuffer(data)
	req, err = http.NewRequest("GET", "/user/check", buf)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(CheckUserBySession)
	handler.ServeHTTP(rr, req)

	expectedStatus = http.StatusNotFound
	if status := rr.Code; status != expectedStatus {
		t.Errorf("Неожиданный код ответа: получено %v, ожидалось %v",
			status, expectedStatus)
	}

	if rr.HeaderMap["Set-Cookies"] != nil {
		t.Errorf("Cookie не были установлены после регистрации!")
	}

	// COOKIE
	cookie := &http.Cookie{
		Name:     "sessionid",
		Value:    "1234567890",
		Expires:  time.Now().AddDate(0, 0, -1),
		Path:     "/",
		HttpOnly: true,
	}
	data, _ = json.Marshal(map[string]string{"nickname": "test", "password": "Test"})
	buf = bytes.NewBuffer(data)
	req, err = http.NewRequest("GET", "/user/check", buf)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(CheckUserBySession)
	req.AddCookie(cookie)
	handler.ServeHTTP(rr, req)
	expectedStatus = http.StatusNotFound
	if status := rr.Code; status != expectedStatus {
		t.Errorf("Неожиданный код ответа: получено %v, ожидалось %v",
			status, expectedStatus)
	}
	if rr.HeaderMap["Set-Cookies"] != nil {
		t.Errorf("Cookie не были установлены после регистрации!")
	}

	cid := helpers.CreateSession("test")
	cookie = &http.Cookie{
		Name:     "sessionid",
		Value:    cid,
		Expires:  time.Now().AddDate(0, 0, -1),
		Path:     "/",
		HttpOnly: true,
	}
	data, _ = json.Marshal(map[string]string{"nickname": "test", "password": "Test"})
	buf = bytes.NewBuffer(data)
	req, err = http.NewRequest("GET", "/user/check", buf)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(CheckUserBySession)
	req.AddCookie(cookie)
	handler.ServeHTTP(rr, req)
	expectedStatus = http.StatusOK
	if status := rr.Code; status != expectedStatus {
		t.Errorf("Неожиданный код ответа: получено %v, ожидалось %v",
			status, expectedStatus)
	}
	if rr.HeaderMap["Set-Cookies"] != nil {
		t.Errorf("Cookie не были установлены после регистрации!")
	}

	data, _ = json.Marshal(map[string]string{"nickname": "test", "password": "Test"})
	buf = bytes.NewBuffer(data)
	req, err = http.NewRequest("GET", "/user", buf)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(GetProfileInfo)
	req.AddCookie(cookie)
	handler.ServeHTTP(rr, req)
	expectedStatus = http.StatusOK
	if status := rr.Code; status != expectedStatus {
		t.Errorf("Неожиданный код ответа: получено %v, ожидалось %v",
			status, expectedStatus)
	}
	if rr.HeaderMap["Set-Cookies"] != nil {
		t.Errorf("Cookie не были установлены после регистрации!")
	}

	data, _ = json.Marshal(map[string]string{"nickname": "test", "password": "Test"})
	buf = bytes.NewBuffer(data)
	req, err = http.NewRequest("POST", "/user/update", buf)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(UpdateProfileInfo)
	// req.AddCookie(cookie)
	handler.ServeHTTP(rr, req)
	expectedStatus = http.StatusNotFound
	if status := rr.Code; status != expectedStatus {
		t.Errorf("Неожиданный код ответа: получено %v, ожидалось %v",
			status, expectedStatus)
	}
	if rr.HeaderMap["Set-Cookies"] != nil {
		t.Errorf("Cookie не были установлены после регистрации!")
	}

	data, _ = json.Marshal(map[string]string{"nickname": "test", "password": "Test"})
	buf = bytes.NewBuffer(data)
	req, err = http.NewRequest("POST", "/user/update", buf)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(UpdateProfileInfo)
	req.AddCookie(cookie)
	handler.ServeHTTP(rr, req)
	expectedStatus = http.StatusOK
	if status := rr.Code; status != expectedStatus {
		t.Errorf("Неожиданный код ответа: получено %v, ожидалось %v",
			status, expectedStatus)
	}
	if rr.HeaderMap["Set-Cookies"] != nil {
		t.Errorf("Cookie не были установлены после регистрации!")
	}

}
