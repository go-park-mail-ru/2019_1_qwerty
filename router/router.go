package router

import (
	"fmt"
	"net/http"
	"os"

	"2019_1_qwerty/api"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Start(port string) error {
	fmt.Println("Api running on port", port)

	var router = mux.NewRouter()
	routerAPI := router.PathPrefix("/api").Subrouter()

	routerAPI.HandleFunc("/user", api.GetProfileInfo).Methods("GET")
	routerAPI.HandleFunc("/user/check", api.CheckUserBySession).Methods("GET")
	routerAPI.HandleFunc("/user/signup", api.CreateUser).Methods("POST", "OPTIONS")
	routerAPI.HandleFunc("/user/create", api.CreateUser).Methods("POST", "OPTIONS")
	routerAPI.HandleFunc("/user/login", api.LoginUser).Methods("POST", "OPTIONS")
	routerAPI.HandleFunc("/user/logout", api.LogoutUser).Methods("GET")
	routerAPI.HandleFunc("/user/update", api.UpdateProfileInfo).Methods("POST", "OPTIONS")
	routerAPI.HandleFunc("/user/avatar", api.UpdateAvatar).Methods("POST", "OPTIONS")

	routerAPI.HandleFunc("/score", api.GetNextAfter).Methods("GET")
	routerAPI.HandleFunc("/score", api.CreateScore).Methods("POST", "OPTIONS")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	origins := handlers.AllowedOrigins([]string{os.Getenv("FRONTEND")})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS", "PUT"})

	return http.ListenAndServe(":"+port, handlers.CORS(origins, headers, methods, handlers.AllowCredentials())(router))
}
