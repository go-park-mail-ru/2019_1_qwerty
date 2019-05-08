package main

import (
	"log"
	"net"
	"os"

	"2019_1_qwerty/auth2/auth"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()

	if err := Open(); err != nil {
		log.Fatal(err.Error())
	}

	auth.RegisterAuthCheckerServer(server, NewSessionManager())

	log.Println("Starting auth server at :" + os.Getenv("PORT"))
	server.Serve(lis)
}
