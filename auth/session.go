package main

import (
	"log"
	"os"

	"2019_1_qwerty/auth/session"

	"github.com/gomodule/redigo/redis"
	"github.com/satori/uuid"

	"golang.org/x/net/context"
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

const sessKeyLen = 10

type SessionManager struct {
	// mu       sync.RWMutex
	sessions map[string]*session.Session
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		// mu:       sync.RWMutex{},
		sessions: map[string]*session.Session{},
	}
}

func (sm *SessionManager) CreateSession(ctx context.Context, in *session.User) (*session.Session, error) {
	sessionID := (uuid.NewV4()).String()
	_, _ = redis.String(RedisConnect.Do("SET", sessionID, in.Nickname, "EX", 86400))
	_, _ = redis.String(RedisConnect.Do("SET", in.Nickname, sessionID, "EX", 86400))
	return &session.Session{ID: sessionID}, nil
}

func (sm *SessionManager) DestroySession(ctx context.Context, in *session.Session) (*session.Status, error) {
	_, _ = RedisConnect.Do("DEL", in.ID)
	user, _ := sm.GetOwner(ctx, in)
	_, _ = RedisConnect.Do("DEL", user.Nickname)
	return &session.Status{Ok: true}, nil
}

func (sm *SessionManager) ValidateSession(ctx context.Context, in *session.Session) (*session.Status, error) {
	_, err := redis.String(RedisConnect.Do("GET", in.ID))
	return &session.Status{Ok: (err != redis.ErrNil)}, nil
}

func (sm *SessionManager) GetOwner(ctx context.Context, in *session.Session) (*session.User, error) {
	res := ""
	str := in.ID[:3]

	if str == "KEY" {
		name := in.ID[3:]
		res, _ = redis.String(RedisConnect.Do("GET", name))
	} else {
		res, _ = redis.String(RedisConnect.Do("GET", in.ID))
	}

	return &session.User{Nickname: res}, nil
}
