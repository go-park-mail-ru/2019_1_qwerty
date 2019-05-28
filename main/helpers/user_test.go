package helpers

import (
	"2019_1_qwerty/main/database"
	"2019_1_qwerty/main/models"
	"database/sql"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const createUserTable = `
CREATE TABLE IF NOT EXISTS users (
    nickname        CITEXT PRIMARY KEY,
    email           CITEXT UNIQUE,
    "password"  text,
    avatar          TEXT DEFAULT 'default0.jpg'
);`

func init() {
	var err error
	database.Database, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := database.Database.Exec(createUserTable); err != nil {
		log.Println(err)
	}
	log.Println("db: Connection: Initialized")
}

func TestUser(t *testing.T) {
	user := models.User{}
	user.Nickname = "Nickname"
	user.Password = "Password"
	user.Email = "mail@mail.mail"
	err := DBUserCreate(&user)
	if err != nil {
		t.Error("Не удалось создать пользователя", err)
	}
	err = DBUserCreate(&user)
	if err == nil {
		t.Error("Удалось создать пользователя с сущ. nickname/email", err)
	}
	user.Email = "m2345@6789mail2"
	err = DBUserUpdate(user.Nickname, &user)
	if err != nil {
		t.Error("Не удалось обновить данные пользователя", err)
	}
	user.Avatar = "test.jpg"
	err = DBUserUpdateAvatar(user.Nickname, user.Avatar)
	if err != nil {
		t.Error("Не удалось обновить аватар пользователя", err)
	}
	err = DBUserValidate(&user)
	if err != nil {
		t.Error("Не удалось провалидировать ПРАВИЛЬНЫЕ данные пользователя", err)
	}
	user.Password = "Password2"
	err = DBUserValidate(&user)
	if err == nil {
		t.Error("НЕПРАВИЛЬНЫЕ данные пароля пользователя прошли валидацию", err)
	}
	user.Password = "Password2"
	user.Nickname = user.Nickname + "test"
	err = DBUserValidate(&user)
	if err == nil {
		t.Error("НЕПРАВИЛЬНЫЕ данные логина пользователя прошли валидацию", err)
	}
	user.Nickname = "Nickname"
	_, err = DBUserGet(user.Nickname)
	if err != nil {
		t.Error("Сущ пользователь не найден", err)
	}
	_, err = DBUserGet(user.Nickname + "wlkfw")
	if err == nil {
		t.Error("Несущ пользователь найден", err)
	}
}
