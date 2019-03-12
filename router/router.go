package Router

import (
	"net/http"

	"../api"
	"github.com/gorilla/mux"
)

// Instance - Экспортируемый экземпляр роутера
var Router = mux.NewRouter()

func init() {
	Router.HandleFunc("/api/user", api.GetProfileInfo).Methods("GET")
	Router.HandleFunc("/api/user/check", api.CheckSession).Methods("GET")
	Router.HandleFunc("/api/user/signup", api.CreateSession).Methods("POST", "OPTIONS")
	Router.HandleFunc("/api/user/login", api.GetSession).Methods("POST", "OPTIONS")
	Router.HandleFunc("/api/user/logout", api.DestroySession).Methods("GET")
	Router.HandleFunc("/api/user/update", api.UpdateProfileInfo).Methods("POST", "OPTIONS")
	Router.HandleFunc("/api/user/avatar", api.UpdateAvatar).Methods("POST", "OPTIONS")

	Router.HandleFunc("/api/score", api.GetNextAfter).Methods("GET")
	Router.HandleFunc("/api/score", api.CreateScore).Methods("POST", "OPTIONS")

	Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
}
