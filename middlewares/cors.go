package middlewares

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

//CORSMiddleware - middleware used for CORS
func CORSMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND"))
		// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, HEAD, DELETE")
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// w.Header().Set("Access-Control-Allow-Credentials", "true")
		headers := handlers.AllowedHeaders([]string{"Content-Type"})
		origins := handlers.AllowedOrigins([]string{os.Getenv("FRONTEND")})
		methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS", "PUT", "DELETE"})
		handler = handlers.CORS(origins, headers, methods, handlers.AllowCredentials())(handler)

		if r.Method == "OPTIONS" {
			return
		}

		handler.ServeHTTP(w, r)
	})
}
