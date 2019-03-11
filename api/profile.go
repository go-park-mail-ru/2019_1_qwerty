package api

import (
	"encoding/json"
        "os"
	"io"
	"net/http"
        "crypto/md5"
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

        if user.Name == "" {
                w.WriteHeader(http.StatusNotFound)
		return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(user)
}

//UploadAvatar - upload avatar to static folder
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
        r.ParseMultipartForm(10 * 1024 * 1024)
        avatar, _, err := r.FormFile("INSERT KEY HERE")
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

        w.WriteHeader(http.StatusOK)
}

//UpdateProfileInfo - updates player data
func UpdateProfileInfo(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("sessionid")

        if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

        user := models.Sessions[string(cookie.Value)]

        newEmail := r.FormValue("INSERT_MAIL")

        if newEmail != "" {
                user.Email = newEmail
        }

        newPassword := r.FormValue("INSERT_PASSWORD")

        if newPassword != "" {
                user.Password = newPassword
        }

        w.WriteHeader(http.StatusOK)
}
