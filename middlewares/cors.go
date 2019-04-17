package middlewares

import (
	"net/http"
	"os"
)

//CORSMiddleware - middleware used for CORS
func CORSMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND"))
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, HEAD")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// headers := handlers.AllowedHeaders([]string{"Content-Type"})
		// origins := handlers.AllowedOrigins([]string{os.Getenv("FRONTEND")})
		// methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS", "PUT"})
		// handler = handlers.CORS(origins, headers, methods, handlers.AllowCredentials())(handler)
		handler.ServeHTTP(w, r)
	})
}
