package tests

import (
	"2019_1_qwerty/handlers"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserCreate(t *testing.T) {
	data, err := json.Marshal(map[string]string{"nickname": "test", "email": "Test@test.ru", "password": "Test"})
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", "/user/create", buf)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateUser)

	handler.ServeHTTP(rr, req)

	expectedStatus := http.StatusCreated
	if status := rr.Code; status != expectedStatus {
		t.Errorf("Неожиданный код ответа: получено %v, ожидалось %v",
			status, expectedStatus)
	}

	if rr.HeaderMap["Set-Cookies"] != nil {
		t.Errorf("Cookie не были установлены после регистрации!")
	}
}
