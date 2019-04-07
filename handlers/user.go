package api

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"2019_1_qwerty/models"
	"2019_1_qwerty/sessions"
)

func init() {
	models.Users = map[string]models.User{}
	models.Sessions = map[string]models.User{}
}

//CreateUser - create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userStruct models.UserRegistration
	if err := json.NewDecoder(r.Body).Decode(&userStruct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userStruct.Name == models.Users[userStruct.Name].Name {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userStruct.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	models.Users[userStruct.Name] = models.User{
		Name:     userStruct.Name,
		Email:    userStruct.Email,
		Password: hashedPassword,
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

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(userStruct.Password))

	if !ok || err != nil {
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

//GetProfileInfo - return player data
func GetProfileInfo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sessionid")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user := models.Sessions[string(cookie.Value)]

	userInfo := models.UserProfile{
		Name:   user.Name,
		Email:  user.Email,
		Score:  user.Score,
		Avatar: user.Avatar,
	}

	if user.Name == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userInfo)
}

//UpdateAvatar - upload avatar to static folder
func UpdateAvatar(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 * 1024 * 1024)
	avatar, _, err := r.FormFile("file")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	defer avatar.Close()

	cookie, err := r.Cookie("sessionid")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user := models.Sessions[string(cookie.Value)]

	generatedName := make([]byte, 8)
	rand.Read(generatedName)
	imageName := fmt.Sprintf("%x", generatedName)

	path := imageName + ".png"

	readyAvatar, _ := os.Create("./static/" + path)
	defer readyAvatar.Close()
	io.Copy(readyAvatar, avatar)

	user.Avatar = path
	models.Sessions[string(cookie.Value)] = user

	currentUser := models.Users[user.Name]
	currentUser.Avatar = path
	models.Users[user.Name] = currentUser

	w.WriteHeader(http.StatusOK)
}

//UpdateProfileInfo - updates player data
func UpdateProfileInfo(w http.ResponseWriter, r *http.Request) {
	var userStruct models.UserChange

	err := json.NewDecoder(r.Body).Decode(&userStruct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("sessionid")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user := models.Sessions[string(cookie.Value)]

	if userStruct.Email != "" {
		user.Email = userStruct.Email
	}

	if userStruct.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userStruct.Password), bcrypt.DefaultCost)

		if err != nil {
			panic(err)
		}

		user.Password = hashedPassword
	}

	models.Sessions[string(cookie.Value)] = user

	currentUser := models.Users[user.Name]
	currentUser.Email = user.Email
	currentUser.Password = user.Password
	models.Users[user.Name] = currentUser

	w.WriteHeader(http.StatusOK)
}
