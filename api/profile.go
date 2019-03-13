package api

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"

	models "../models"
)

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
		user.Password = userStruct.Password
	}

	models.Sessions[string(cookie.Value)] = user

	currentUser := models.Users[user.Name]
	currentUser.Email = user.Email
	currentUser.Password = user.Password
	models.Users[user.Name] = currentUser

	w.WriteHeader(http.StatusOK)
}
