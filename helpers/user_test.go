package helpers

import (
	"2019_1_qwerty/models"
	"database/sql"
	"io/ioutil"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if query, err := ioutil.ReadFile("sql/create.sql"); err != nil {
		log.Println(err)
	} else {
		if _, err := db.Exec(string(query)); err != nil {
			log.Println(err)
		}
	}
	log.Println("db: Connection: Initialized")

}

func TestUser(t *testing.T) {
	// _ = godotenv.Load()
	// if err := database.Open(); err != nil {
	// 	log.Println(err.Error())
	// }
	// database.Close()
	// rc = sqlite3_open("file::memory:", &db)

	user := models.User{}
	user.Nickname = "Nickname"
	user.Password = "Password"
	user.Email = "mail@mail.mail"
	err := db.DBUserCreate(&user)
	if err != nil {
		t.Error("Не удалось создать пользователя", err)
	}
	user.Email = "mail@mail.mail"
	err = DBUserUpdate(user.Nickname, &user)
	if err != nil {
		t.Error("Не удалось обновить данные пользователя", err)
	}
	err = DBUserValidate(&user)
	if err != nil {
		t.Error("Не удалось провалидировать ПРАВИЛЬНЫЕ данные пользователя", err)
	}
	user.Password = "Password2"
	err = DBUserValidate(&user)
	if err != nil {
		t.Error("НЕПРАВИЛЬНЫЕ данные пользователя прошли валидацию", err)
	}
	user2, _ := DBUserGet(user.Nickname)
	if (user.Nickname != user2.Nickname) || (user.Email != user2.Email) {
		t.Error("Данные на пользователя не совпадают", err)
	}
}
