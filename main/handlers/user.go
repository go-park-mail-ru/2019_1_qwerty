package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"2019_1_qwerty/helpers"
	"2019_1_qwerty/models"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go"
	"github.com/satori/uuid"
)

var s3Client *minio.Client
var endpoint string

const bucketName = "qwertys3"

func init() {
	var err error
	_ = godotenv.Load()
	ssl, _ := strconv.ParseBool(os.Getenv("S3_SSL"))
	location := os.Getenv("S3_LOCATION")
	secretAccessKey := os.Getenv("S3_SECRETACCESSKEY")
	accessKeyID := os.Getenv("S3_ACCESSKEYID")
	endpoint = os.Getenv("S3_ENDPOINT")
	s3Client, err = minio.New(endpoint, accessKeyID, secretAccessKey, ssl)
	if err != nil {
		log.Fatalln(err)
	}

	err = s3Client.MakeBucket(bucketName, location)
	if err != nil {
		exists, err := s3Client.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s bucket\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s bucket\n", bucketName)
	}
}

//UpdateAvatar - upload avatar to static folder
func UpdateAvatar(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 * 1024 * 1024)
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		Hits.WithLabelValues(string(http.StatusBadRequest), r.URL.String()).Inc()
		FooCount.Add(1)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	cookie, err := r.Cookie("sessionid")
	if err != nil {
		Hits.WithLabelValues(string(http.StatusNotFound), r.URL.String()).Inc()
		FooCount.Add(1)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	nickname := helpers.GetOwner(string(cookie.Value))

	objectName := (uuid.NewV4()).String() + filepath.Ext(fileHeader.Filename)

	_, err = s3Client.PutObject(bucketName, objectName, file, -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		log.Fatalln(err)
	}

	err = helpers.DBUserUpdateAvatar(nickname, objectName)
	if err != nil {
		Hits.WithLabelValues(string(http.StatusNotFound), r.URL.String()).Inc()
		FooCount.Add(1)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	Hits.WithLabelValues(string(http.StatusOK), r.URL.String()).Inc()
	FooCount.Add(1)
	w.WriteHeader(http.StatusOK)
}

// CreateUser - Создание пользователя
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	_ = json.NewDecoder(r.Body).Decode(&user)
	err := helpers.DBUserCreate(&user)
	if err != nil {
		// Если логин и пароль совпадут, то пользователь будет залогинен, а не зареган
		if !helpers.LoginUser(user.Nickname, user.Password) {
			Hits.WithLabelValues(string(http.StatusConflict), r.URL.String()).Inc()
			FooCount.Add(1)
			w.WriteHeader(http.StatusConflict)
			return
		}
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "sessionid",
		Value:    helpers.CreateSession(user.Nickname),
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	})
	Hits.WithLabelValues(string(http.StatusCreated), r.URL.String()).Inc()
	FooCount.Add(1)
	w.WriteHeader(http.StatusCreated)
}

// LoginUser - авторизация
func LoginUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	_ = json.NewDecoder(r.Body).Decode(&user)

	if !helpers.LoginUser(user.Nickname, user.Password) {
		Hits.WithLabelValues(string(http.StatusNotFound), r.URL.String()).Inc()
		FooCount.Add(1)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "sessionid",
		Value:    helpers.CreateSession(user.Nickname),
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	})
	Hits.WithLabelValues(string(http.StatusOK), r.URL.String()).Inc()
	FooCount.Add(1)
	w.WriteHeader(http.StatusOK)
}

//CheckUserBySession - user authorization status // Разобрать говно потом
func CheckUserBySession(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("sessionid"); err == nil {
		if !helpers.ValidateSession(string(cookie.Value)) {
			http.SetCookie(w, &http.Cookie{
				Name:     "sessionid",
				Value:    "",
				Expires:  time.Now().AddDate(0, 0, -1),
				Path:     "/",
				HttpOnly: true,
			})
			Hits.WithLabelValues(string(http.StatusNotFound), r.URL.String()).Inc()
			FooCount.Add(1)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		Hits.WithLabelValues(string(http.StatusOK), r.URL.String()).Inc()
		FooCount.Add(1)
		w.WriteHeader(http.StatusOK)
	} else {
		Hits.WithLabelValues(string(http.StatusNotFound), r.URL.String()).Inc()
		FooCount.Add(1)
		w.WriteHeader(http.StatusNotFound)
	}
}

//LogoutUser - deauthorization // Разобрать говно потом
func LogoutUser(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("sessionid"); err == nil {
		http.SetCookie(w, &http.Cookie{
			Name:     "sessionid",
			Value:    "",
			Expires:  time.Now().AddDate(0, 0, -1),
			Path:     "/",
			HttpOnly: true,
		})
		helpers.DestroySession(string(cookie.Value))
	}
	Hits.WithLabelValues(string(http.StatusOK), r.URL.String()).Inc()
	FooCount.Add(1)
	w.WriteHeader(http.StatusOK)
}

//GetProfileInfo - return player data
func GetProfileInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cookie, err := r.Cookie("sessionid")
	if err != nil {
		Hits.WithLabelValues(string(http.StatusNotFound), r.URL.String()).Inc()
		FooCount.Add(1)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	user := helpers.GetOwner(string(cookie.Value))
	res, _ := helpers.DBUserGet(user)

	if res.Avatar != "" {
		reqParams := make(url.Values)
		reqParams.Set("response-content-disposition", "attachment; filename="+res.Avatar)
		presignedURL, err := s3Client.PresignedGetObject(bucketName, res.Avatar, time.Second*24*60*60, reqParams)
		if err != nil {
			log.Println(err)
		}
		res.Avatar = presignedURL.String()
	}
	Hits.WithLabelValues(string(http.StatusOK), r.URL.String()).Inc()
	FooCount.Add(1)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

//UpdateProfileInfo - updates player data
func UpdateProfileInfo(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	_ = json.NewDecoder(r.Body).Decode(&user)
	cookie, err := r.Cookie("sessionid")
	if err != nil {
		Hits.WithLabelValues(string(http.StatusNotFound), r.URL.String()).Inc()
		FooCount.Add(1)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	nickname := helpers.GetOwner(string(cookie.Value))
	err = helpers.DBUserUpdate(nickname, &user)
	if err != nil {
		Hits.WithLabelValues(string(http.StatusNotFound), r.URL.String()).Inc()
		FooCount.Add(1)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	Hits.WithLabelValues(string(http.StatusOK), r.URL.String()).Inc()
	FooCount.Add(1)
	w.WriteHeader(http.StatusOK)
}
