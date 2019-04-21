package middlewares

import (
	"net/http"
)

//LogMiddleware - [FOR LOCAL TESTS ONLY] middleware for logs/info
func LogMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// begin := time.Now().Format("2006-01-02 15:04:05")
		// log.Println(begin, ":", r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}
