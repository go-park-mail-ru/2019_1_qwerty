package api

import (
	"encoding/json"
	"net/http"
	"time"

	models "../models"
	uuid "github.com/satori/uuid"
)

func init() {
	models.Users = map[string]models.User{}
	models.Sessions = map[string]models.User{}
}

//CreateSession - create user
func CreateSession(w http.ResponseWriter, r *http.Request) {
	var userStruct models.UserRegistration
	err := json.NewDecoder(r.Body).Decode(&userStruct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userStruct.Name == models.Users[userStruct.Name].Name {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	models.Users[userStruct.Name] = models.User{
		Name:     userStruct.Name,
		Email:    userStruct.Email,
		Password: userStruct.Password,
		Score:    0,
		Avatar:   "default.jpg",
	}

	sessionID, _ := uuid.NewV4()
	models.Sessions[sessionID.String()] = models.Users[userStruct.Name]

	http.SetCookie(w, &http.Cookie{
		Name:     "sessionid",
		Value:    sessionID.String(),
		Expires:  time.Now().Add(60 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusCreated)
}

//GetSession - authorization
func GetSession(w http.ResponseWriter, r *http.Request) {
	var userStruct models.UserLogin
	err := json.NewDecoder(r.Body).Decode(&userStruct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, ok := models.Users[userStruct.Name]

	if !ok || userStruct.Password != user.Password {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	sessionID, _ := uuid.NewV4()
	models.Sessions[sessionID.String()] = user

	http.SetCookie(w, &http.Cookie{
		Name:     "sessionid",
		Value:    sessionID.String(),
		Expires:  time.Now().Add(60 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusOK)
}

//CheckSession - user authorization status
func CheckSession(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("sessionid")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//DestroySession - deauthorization
func DestroySession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sessionid")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "sessionid",
		Value:    "",
		Expires:  time.Now().AddDate(0, 0, -1),
		Path:     "/",
		HttpOnly: true,
	})

	delete(models.Sessions, string(cookie.Value))
	w.WriteHeader(http.StatusOK)
}
