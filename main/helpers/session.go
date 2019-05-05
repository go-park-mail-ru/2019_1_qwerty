package helpers

import (
	"context"
	"log"

	"2019_1_qwerty/helpers/session"

	"google.golang.org/grpc"
)

var (
	grcpConn    *grpc.ClientConn
	sessManager session.AuthCheckerClient
)

func Open() error {
	var err error
	grcpConn, err = grpc.Dial(
		"backend_auth:8080",
		grpc.WithInsecure(),
	)
	if err != nil {
		return err
	}
	sessManager = session.NewAuthCheckerClient(grcpConn)
	return nil
}

func CreateSession(user string) string {
	ctx := context.Background()
	sessId, err := sessManager.CreateSession(ctx,
		&session.User{
			Nickname: user,
		})
	if err != nil {
		log.Println(err)
	}
	return sessId.ID
}

func DestroySession(sessionID string) {
	ctx := context.Background()
	sessManager.DestroySession(ctx,
		&session.Session{
			ID: sessionID,
		})
}

func ValidateSession(sessionID string) bool {
	ctx := context.Background()
	status, err := sessManager.ValidateSession(ctx,
		&session.Session{
			ID: sessionID,
		})
	if err != nil {
		log.Println(err)
	}
	return status.Ok
}

func GetOwner(sessionID string) string {
	ctx := context.Background()
	user, err := sessManager.GetOwner(ctx,
		&session.Session{
			ID: sessionID,
		})
	if err != nil {
		log.Println(err)
	}
	return user.Nickname
}
