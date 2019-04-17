package router

import (
	api "2019_1_qwerty/handlers"
	"2019_1_qwerty/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Start - router logic
func Start(port string) error {
	log.Println("Api running on port", port)

	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	routerAPI := router.PathPrefix("/api").Subrouter()
	routerAPI.HandleFunc("/user/signup", api.CreateUser).Methods("POST", "OPTIONS")
	routerAPI.HandleFunc("/user/create", api.CreateUser).Methods("POST", "OPTIONS")
	routerAPI.HandleFunc("/user/login", api.LoginUser).Methods("POST", "OPTIONS")
	routerAPI.HandleFunc("/score", api.GetNextAfter).Methods("GET")
	routerAPI.HandleFunc("/score", api.CreateScore).Methods("POST", "OPTIONS")
	routerAPI.HandleFunc("/user/check", api.CheckUserBySession).Methods("GET")
	routerAPI.HandleFunc("/ws", api.WebsocketConn).Methods("GET", "POST")

	routerLogged := router.PathPrefix("/api").Subrouter()
	routerLogged.Use(middlewares.AuthorizationMiddleware)
	routerLogged.HandleFunc("/user", api.GetProfileInfo).Methods("GET", "OPTIONS")
	routerLogged.HandleFunc("/user/update", api.UpdateProfileInfo).Methods("POST", "OPTIONS")
	routerLogged.HandleFunc("/user/avatar", api.UpdateAvatar).Methods("POST", "OPTIONS")
	routerLogged.HandleFunc("/user/logout", api.LogoutUser).Methods("GET")

	serverHandler := middlewares.LogMiddleware(middlewares.ErrorMiddleware(middlewares.CORSMiddleware(router)))

	return http.ListenAndServe(":"+port, serverHandler)
}
