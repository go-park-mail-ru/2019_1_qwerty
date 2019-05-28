package router

import (
	api "2019_1_qwerty/handlers"
	"2019_1_qwerty/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//Start - router logic
func Start(port string) error {
	log.Println("Api running on port", port)

	router := mux.NewRouter()
	// router.HandleFunc("/metrics", promhttp.Handler)
	prometheus.MustRegister(api.Hits)
	prometheus.MustRegister(api.FooCount)
	router.Handle("/metrics", promhttp.Handler())
	routerAPI := router.PathPrefix("/api").Subrouter()
	routerAPI.HandleFunc("/user/signup", api.CreateUser).Methods("POST", "OPTIONS")
	routerAPI.HandleFunc("/user/create", api.CreateUser).Methods("POST", "OPTIONS")
	routerAPI.HandleFunc("/user/login", api.LoginUser).Methods("POST", "OPTIONS")
	routerAPI.HandleFunc("/score", api.GetNextAfter).Methods("GET")
	routerAPI.HandleFunc("/user/check", api.CheckUserBySession).Methods("GET")
	routerAPI.HandleFunc("/ws", api.WebsocketConn).Methods("GET", "POST", "OPTIONS")

	routerLogged := router.PathPrefix("/api").Subrouter()
	routerLogged.Use(middlewares.AuthorizationMiddleware)
	routerLogged.HandleFunc("/user", api.GetProfileInfo).Methods("GET", "OPTIONS")
	routerLogged.HandleFunc("/user/update", api.UpdateProfileInfo).Methods("POST", "OPTIONS")
	routerLogged.HandleFunc("/user/avatar", api.UpdateAvatar).Methods("POST", "OPTIONS")
	routerLogged.HandleFunc("/user/logout", api.LogoutUser).Methods("GET")

	serverHandler := middlewares.LogMiddleware(middlewares.ErrorMiddleware(middlewares.CORSMiddleware(router)))

	return http.ListenAndServe(":"+port, serverHandler)
}
