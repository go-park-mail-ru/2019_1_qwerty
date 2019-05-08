package main

import (
	"log"
	"os"

	"2019_1_qwerty/database"
	"2019_1_qwerty/router"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	if err := database.Open(); err != nil {
		log.Fatal(err.Error())
	}
	defer database.Close()

	err := router.Start(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
