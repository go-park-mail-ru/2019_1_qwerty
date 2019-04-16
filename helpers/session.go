package helpers

import (
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/uuid"
)

var (
	RedisConnect redis.Conn
)

func Open() error {
	var err error
	RedisConnect, err = redis.DialURL(os.Getenv("REDIS_URL"))
	if err != nil {
		return err
	}
	log.Println("Redis: Connection: Initialized")
	return nil
}

func CreateSession(user string) string {
	sessionID := (uuid.NewV4()).String()
	_, _ = redis.String(RedisConnect.Do("SET", sessionID, user, "EX", 86400))
	return sessionID
}

func DestroySession(sessionID string) {
	_, _ = RedisConnect.Do("DEL", sessionID)
}

func ValidateSession(sessionID string) bool {
	_, err := redis.String(RedisConnect.Do("GET", sessionID))
	return (err != redis.ErrNil)
}

func GetOwner(sessionID string) string {
	res, _ := redis.String(RedisConnect.Do("GET", sessionID))
	return res
}
