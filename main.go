package main

import (
	"fmt"
        "net/http"
        "./api"
        "github.com/gorilla/mux"
        //"github.com/gorilla/handlers"
        "github.com/rs/cors"
)

func main() {
        c := cors.New(cors.Options {
                AllowedOrigins: []string{"http://localhost:8000/"},
                AllowedMethods: []string{"GET, POST, OPTIONS"},
                AllowedHeaders: []string{"Content-Type", "Accept", "Origin"},
                AllowCredentials: true,
                OptionsPassthrough: true,
        })


        router := mux.NewRouter()
        router.HandleFunc("/api/user/check", api.CheckSession).Methods("GET")
        router.HandleFunc("/api/user/signup", api.CreateSession).Methods("OPTIONS, POST")


	fmt.Println("Api running on port 8080")
	http.ListenAndServe(":8080", c.Handler(router))
}
