package tests

import (
	"2019_1_qwerty/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestScoreCreate(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/score?points=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateScore)

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
	handler := http.HandlerFunc(handlers.GetNextAfter)

	handler.ServeHTTP(rr, req)

	expectedStatus := http.StatusOK
	if status := rr.Code; status != expectedStatus {
		t.Errorf("Неожиданный код ответа: получено %v, ожидалось %v",
			status, expectedStatus)
	}
}
