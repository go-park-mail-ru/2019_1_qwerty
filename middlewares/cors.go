package middlewares

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

//CORSMiddleware - middleware used for CORS
func CORSMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := handlers.AllowedHeaders([]string{"Content-Type"})
		origins := handlers.AllowedOrigins([]string{os.Getenv("FRONTEND")})
		methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS", "PUT"})

		handler = handlers.CORS(origins, headers, methods, handlers.AllowCredentials())(handler)

		handler.ServeHTTP(w, r)
	})
}
