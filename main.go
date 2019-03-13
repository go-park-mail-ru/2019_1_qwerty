package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	router "./router"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	origins := handlers.AllowedOrigins([]string{os.Getenv("FRONTEND")})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS", "PUT"})

	fmt.Println("Api running on port 8080")
	http.ListenAndServe(":8080", handlers.CORS(origins, headers, methods, handlers.AllowCredentials())(router.Router))
}
