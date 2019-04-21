package middlewares

import (
	"net/http"
	"os"
	"strings"
)

//CORSMiddleware - middleware used for CORS
func CORSMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		origins := strings.Split(os.Getenv("FRONTEND"), ",")

		for _, v := range origins {

			if origin == v {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}

		}

		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, HEAD, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			return
		}

		handler.ServeHTTP(w, r)
	})
}
