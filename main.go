package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"./api"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/user/check", api.CheckSession).Methods("GET")
	router.HandleFunc("/api/user/signup", api.CreateSession).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/login", api.GetSession).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/logout", api.DestroySession).Methods("GET")

	router.HandleFunc("/api/score", api.GetNextAfter).Methods("GET")
	router.HandleFunc("/api/score", api.CreateScore).Methods("POST", "OPTIONS")


	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	origins := handlers.AllowedOrigins([]string{os.Getenv("FRONTEND")})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS", "PUT"})

	fmt.Println("Api running on port 8080")
	http.ListenAndServe(":8080", handlers.CORS(origins, headers, methods, handlers.AllowCredentials())(router))
}
