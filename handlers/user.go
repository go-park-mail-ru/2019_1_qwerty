package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"2019_1_qwerty/helpers"
	"2019_1_qwerty/models"
)

func init() {
	models.Sessions = map[string]string{}
}

// CreateUser - Создание пользователя
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("encode: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := helpers.DBUserCreate(&user)
	if err != nil {
		log.Println("CreateUserm: DBUserCreate: ", err)
		w.WriteHeader(http.StatusConflict)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "sessionid",
		Value:    helpers.CreateSession(user.Nickname),
		Expires:  time.Now().Add(60 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusCreated)
}

// LoginUser - авторизация
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = helpers.DBUserValidate(&user)
	if err != nil {
		log.Println("LoginUser: DBValidateUser: ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "sessionid",
		Value:    helpers.CreateSession(user.Nickname),
		Expires:  time.Now().Add(60 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusOK)
}

//CheckUserBySession - user authorization status // Разобрать говно потом
func CheckUserBySession(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("sessionid"); err == nil {
		w.WriteHeader(http.StatusOK)
		// if helpers.ValidateSession(cookie.String()) {
		// 	w.WriteHeader(http.StatusOK)
		// } else {
		// 	w.WriteHeader(http.StatusNotFound)
		// }
	} else {
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
	w.WriteHeader(http.StatusOK)
}

//GetProfileInfo - return player data
func GetProfileInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cookie, err := r.Cookie("sessionid")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	user := helpers.GetOwner(string(cookie.Value))
	res, err := helpers.DBUserGet(user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	// helpers.ErroRouter(&w, res, err, http.StatusOK)
}

// //UpdateAvatar - upload avatar to static folder
// func UpdateAvatar(w http.ResponseWriter, r *http.Request) {
// 	r.ParseMultipartForm(5 * 1024 * 1024)
// 	avatar, _, err := r.FormFile("file")
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}
// 	defer avatar.Close()

// 	cookie, err := r.Cookie("sessionid")
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// 	user := models.Sessions[string(cookie.Value)]

// 	generatedName := make([]byte, 8)
// 	rand.Read(generatedName)
// 	imageName := fmt.Sprintf("%x", generatedName)

// 	path := imageName + ".png"

// 	readyAvatar, _ := os.Create("./static/" + path)
// 	defer readyAvatar.Close()
// 	io.Copy(readyAvatar, avatar)

// 	user.Avatar = path // имя аватарки
// 	// models.Sessions[string(cookie.Value)] = user

// 	currentUser := models.Users[user.Name]
// 	currentUser.Avatar = path
// 	models.Users[user.Name] = currentUser

// 	w.WriteHeader(http.StatusOK)
// }

// //UpdateProfileInfo - updates player data
// func UpdateProfileInfo(w http.ResponseWriter, r *http.Request) {
// 	var userStruct models.UserChange

// 	err := json.NewDecoder(r.Body).Decode(&userStruct)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	cookie, err := r.Cookie("sessionid")
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// 	user := models.Sessions[string(cookie.Value)]
// 	res, err = DBUpdateUser(nickname)

// 	models.Sessions[string(cookie.Value)] = user

// 	currentUser := models.Users[user.Name]
// 	currentUser.Email = user.Email
// 	currentUser.Password = user.Password
// 	models.Users[user.Name] = currentUser

// 	w.WriteHeader(http.StatusOK)
// }
