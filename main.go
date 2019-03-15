package main

import (
	"log"
	"os"

	"./router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = router.Start(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}
