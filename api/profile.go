package api

import (
	models "../models"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
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

	fmt.Println(user)

	path := user.Name + strconv.Itoa(rand.Intn(1000)) + ".png"

	readyAvatar, _ := os.Create("./static/" + path)
	defer readyAvatar.Close()
	io.Copy(readyAvatar, avatar)

	user.Avatar = path
	models.Sessions[string(cookie.Value)] = user

	fmt.Println(user)

	w.WriteHeader(http.StatusOK)
}

//UpdateProfileInfo - updates player data
func UpdateProfileInfo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cookie, err := r.Cookie("sessionid")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	newUserStruct := models.UserChange{}

	jsonErr := json.Unmarshal(body, &newUserStruct)

	if jsonErr != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user := models.Sessions[string(cookie.Value)]

	if newUserStruct.Email != "" {
		user.Email = newUserStruct.Email
	}

	if newUserStruct.Password != "" {
		user.Password = newUserStruct.Password
	}

	models.Sessions[string(cookie.Value)] = user
	currentUser := models.Users[user.Name]
	currentUser.Email = user.Email
	currentUser.Password = user.Password
	models.Users[user.Name] = currentUser

	w.WriteHeader(http.StatusOK)
}
