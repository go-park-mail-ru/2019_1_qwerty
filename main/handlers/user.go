package handlers

import (
	"io/ioutil"
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
	log.Println(r.Header.Get("Content-Type"))
	err := r.ParseMultipartForm(5 * 1024 * 1025)
	if err != nil {
		log.Println("Ошибка при парсинге тела запроса", err)
		ErrorMux(&w, r, http.StatusInternalServerError)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		log.Println("Ошибка при получении файла из тела запроса", err)
		ErrorMux(&w, r, http.StatusInternalServerError)
		return
	}
	defer file.Close()

	cookie, err := r.Cookie("sessionid")

	if err != nil {
		ErrorMux(&w, r, http.StatusNotFound)
		return
	}

	nickname := helpers.GetOwner(string(cookie.Value))

	objectName := (uuid.NewV4()).String() + filepath.Ext(fileHeader.Filename)
	_, err = s3Client.PutObject(bucketName, objectName, file, -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})

	if err != nil {
		log.Println(err)
		ErrorMux(&w, r, http.StatusInternalServerError)
	}

	err = helpers.DBUserUpdateAvatar(nickname, objectName)

	if err != nil {
		ErrorMux(&w, r, http.StatusNotFound)
		return
	}

	ErrorMux(&w, r, http.StatusOK)
}

// CreateUser - Создание пользователя
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	user := models.User{}
	user.UnmarshalJSON(body)
	err := helpers.DBUserCreate(&user)

	if err != nil {
		if !helpers.LoginUser(user.Nickname, user.Password) {
			ErrorMux(&w, r, http.StatusConflict)
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

	ErrorMux(&w, r, http.StatusCreated)
}

// LoginUser - авторизация
func LoginUser(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	user := models.User{}
	user.UnmarshalJSON(body)

	if !helpers.LoginUser(user.Nickname, user.Password) {
		ErrorMux(&w, r, http.StatusNotFound)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "sessionid",
		Value:    helpers.CreateSession(user.Nickname),
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	})

	ErrorMux(&w, r, http.StatusOK)
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
			ErrorMux(&w, r, http.StatusNotFound)
			return
		}
		ErrorMux(&w, r, http.StatusOK)
	} else {
		ErrorMux(&w, r, http.StatusNotFound)
	}
}

//LogoutUser - deauthorization // Разобрать говно потом
func LogoutUser(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("sessionid"); err == nil {
		helpers.DestroySession(string(cookie.Value))
		http.SetCookie(w, &http.Cookie{
			Name:     "sessionid",
			Value:    "",
			Expires:  time.Now().AddDate(0, 0, -1),
			Path:     "/",
			HttpOnly: true,
		})
	}
	ErrorMux(&w, r, http.StatusOK)
}

//GetProfileInfo - return player data
func GetProfileInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cookie, err := r.Cookie("sessionid")

	if err != nil {
		ErrorMux(&w, r, http.StatusNotFound)
		return
	}

	user := helpers.GetOwner(string(cookie.Value))
	res, _ := helpers.DBUserGet(user)
	res.Score, _ = helpers.DBUserGetScore(res.Nickname)

	if res.Avatar != "" {
		reqParams := make(url.Values)
		reqParams.Set("response-content-disposition", "attachment; filename="+res.Avatar)
		presignedURL, err := s3Client.PresignedGetObject(bucketName, res.Avatar, time.Second*24*60*60, reqParams)
		if err != nil {
			log.Println(err)
		}
		url := presignedURL.String()
		url = url[:4] + "s" + url[4:]
		res.Avatar = url
	}

	ErrorMux(&w, r, http.StatusOK)
	result, _ := res.MarshalJSON()
	w.Write(result)
}

//UpdateProfileInfo - updates player data
func UpdateProfileInfo(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	user := models.User{}
	user.UnmarshalJSON(body)
	cookie, err := r.Cookie("sessionid")

	if err != nil {
		ErrorMux(&w, r, http.StatusNotFound)
		return
	}

	nickname := helpers.GetOwner(string(cookie.Value))
	err = helpers.DBUserUpdate(nickname, &user)
	if err != nil {
		log.Println(err)
		ErrorMux(&w, r, http.StatusNotFound)
		return
	}

	ErrorMux(&w, r, http.StatusOK)
}
