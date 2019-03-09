package main

import (
	"fmt"
        "net/http"
        "./api"
        "github.com/gorilla/mux"
        "github.com/gorilla/handlers"
)

func main() {
        router := mux.NewRouter()
        router.HandleFunc("/api/user/check", api.CheckSession).Methods("GET")
        router.HandleFunc("/api/user/signup", api.CreateSession).Methods("POST", "OPTIONS")
        router.HandleFunc("/api/user/login", api.GetSession).Methods("POST", "OPTIONS")
        router.HandleFunc("/api/user/logout", api.DestroySession).Methods("GET")

        headers := handlers.AllowedHeaders([]string{"Content-Type"})
        origins := handlers.AllowedOrigins([]string{"http://localhost:8000"})
        methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS", "PUT"})

	fmt.Println("Api running on port 8080")
	http.ListenAndServe(":8080", handlers.CORS(origins, headers, methods, handlers.AllowCredentials())(router))
}
