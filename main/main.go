package main

import (
	"log"
	"os"

	"2019_1_qwerty/main/database"
	"2019_1_qwerty/main/helpers"
	"2019_1_qwerty/main/router"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	if err := database.Open(); err != nil {
		log.Fatal(err.Error())
	}
	defer database.Close()
	if err := helpers.Open(); err != nil {
		log.Fatal(err.Error())
	}

	if err := helpers.OpenAuth(); err != nil {
		log.Fatal(err.Error())
	}

	err := router.Start(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
