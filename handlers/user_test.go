package handlers

import (
	_ "github.com/mattn/go-sqlite3"
)

// func TestUserCreate(t *testing.T) {
// 	// var err error
// 	db, err := sql.Open("sqlite3", ":memory")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	data, _ := json.Marshal(map[string]string{"nickname": "test", "email": "Test@test.ru", "password": "Test"})
// 	buf := bytes.NewBuffer(data)
// 	req, err := http.NewRequest("POST", "/user/create", buf)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(CreateUser)

// 	handler.ServeHTTP(rr, req)

// expectedStatus := http.StatusCreated
// if status := rr.Code; status != expectedStatus {
// 	t.Errorf("Неожиданный код ответа: получено %v, ожидалось %v",
// 		status, expectedStatus)
// }

// if rr.HeaderMap["Set-Cookies"] != nil {
// 	t.Errorf("Cookie не были установлены после регистрации!")
// }
// }
