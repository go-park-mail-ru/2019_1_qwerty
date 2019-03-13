package api

import (
	"encoding/json"
	"net/http"
	"time"

	models "../models"
	sessions "../sessions"
)

func init() {
	models.Users = map[string]models.User{}
	models.Sessions = map[string]models.User{}
}

//CreateUser - create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	http.SetCookie(w, &http.Cookie{
		Name:     "sessionid",
		Value:    sessions.CreateSession(userStruct.Name),
		Expires:  time.Now().Add(60 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusCreated)
}

//LoginUser - authorization
func LoginUser(w http.ResponseWriter, r *http.Request) {
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

	http.SetCookie(w, &http.Cookie{
		Name:     "sessionid",
		Value:    sessions.CreateSession(user.Name),
		Expires:  time.Now().Add(60 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusOK)
}

//CheckUserBySession - user authorization status
func CheckUserBySession(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("sessionid"); err == nil {
		w.WriteHeader(http.StatusOK)
		// if sessions.ValidateSession(cookie.String()) {
		// 	w.WriteHeader(http.StatusOK)
		// } else {
		// 	w.WriteHeader(http.StatusNotFound)
		// }
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

//LogoutUser - deauthorization
func LogoutUser(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("sessionid"); err == nil {
		http.SetCookie(w, &http.Cookie{
			Name:     "sessionid",
			Value:    "",
			Expires:  time.Now().AddDate(0, 0, -1),
			Path:     "/",
			HttpOnly: true,
		})
		sessions.DestroySession(string(cookie.Value))
	}
	w.WriteHeader(http.StatusOK)
}
