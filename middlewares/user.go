package middlewares

import (
	"fmt"
	"net/http"
)

//AuthorizationMiddleware - middleware to check if user is logged in
func AuthorizationMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("auth check")
		_, err := r.Cookie("sessionid")

		if err != nil {
			fmt.Println("user is not logged in!", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusFound)
			return
		}

		fmt.Println("logged!")

		handler.ServeHTTP(w, r)
	})
}
