package api

import (
	"encoding/json"
        "os"
	"io"
        "io/ioutil"
	"net/http"
        "crypto/md5"
        "fmt"
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

        userInfo := models.UserProfile {
                Name: user.Name,
                Email: user.Email,
                Score: user.Score,
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

//UploadAvatar - upload avatar to static folder
func UploadAvatar(w http.ResponseWriter, r *http.Request) {

        fmt.Println("UploadAvatar")
        fmt.Println(r.Body)

        r.ParseMultipartForm(10 * 1024 * 1024)
        avatar, _, err := r.FormFile("file")
        defer avatar.Close()

        if err != nil {
                w.WriteHeader(http.StatusNotFound)
                return
        }

        avatarName := md5.New()
        io.Copy(avatarName, avatar)
        path := string(avatarName.Sum(nil)) + ".jpg"
        image, _ := os.Create("../static/" + path)
        defer image.Close()
        io.Copy(image, avatar)

        cookie, err := r.Cookie("sessionid")

        if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

        user := models.Sessions[string(cookie.Value)]
        user.Avatar = path
        models.Sessions[string(cookie.Value)] = user

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
