package helpers

import (
	"context"
	"log"

	"2019_1_qwerty/main/helpers/auth"

	"google.golang.org/grpc"
)

var (
	grcpConn2    *grpc.ClientConn
	sessManager2 auth.AuthCheckerClient
)

func OpenAuth() error {
	var err error
	grcpConn2, err = grpc.Dial(
		"backend_auth2:8080",
		grpc.WithInsecure(),
	)
	if err != nil {
		return err
	}
	sessManager2 = auth.NewAuthCheckerClient(grcpConn2)
	return nil
}

// бестолковая обёртка над удалённой функцией из-за соединения
func LoginUser(user string, password string) bool {
	ctx := context.Background()
	status, err := sessManager2.LoginUser(ctx,
		&auth.User{
			Nickname: user,
			Password: password,
		})
	if err != nil {
		log.Println(err)
	}
	return status.Ok
}
