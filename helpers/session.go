package helpers

import (
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
	uuid "github.com/satori/uuid"
)

var (
	c redis.Conn
)

func init() {
	_ = godotenv.Load()
	var err error
	c, err = redis.DialURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Println("Redis: ", err)
		return
	}
	log.Println("Redis: Connection: Initialized")
}

func CreateSession(user string) string {
	sessionID := (uuid.NewV4()).String()
	_, _ := redis.String(c.Do("SET", sessionID, user, "EX", 86400))
	return sessionID
}

func DestroySession(sessionID string) {
	_, _ := c.Do("DEL", sessionID)
}

func ValidateSession(sessionID string) bool {
	_, err := redis.String(c.Do("GET", sessionID))
	if err == redis.ErrNil {
		return false
	} else {
		return true
	}
}

func GetOwner(sessionID string) string {
	res, _ := redis.String(c.Do("GET", sessionID))
	return res
}
