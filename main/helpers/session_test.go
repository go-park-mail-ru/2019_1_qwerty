package helpers

import (
	"log"
	"testing"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

func init() {
	var err error
	_ = godotenv.Load()
	RedisConnect, err = redis.DialURL("redis://redis@localhost:6379/0")
	if err != nil {
		log.Println(err)
	}
	log.Println("Redis: Connection: Initialized")
}
func TestSession(t *testing.T) {
	nickname := "test_user"
	sid := CreateSession(nickname)
	owner := GetOwner(sid)
	if owner != nickname {
		t.Error("Не удалось создать сессию пользователя")
	}
	ok := ValidateSession(sid)
	if !ok {
		t.Error("Cессия пользователя не валидна")
	}
	DestroySession(sid)
	ok = ValidateSession(sid)
	if ok {
		t.Error("Cессия пользователя не была удалена")
	}
}
