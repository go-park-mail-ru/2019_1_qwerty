package middlewares

import (
	"2019_1_qwerty/main/helpers"
	"net/http"
)

//AuthorizationMiddleware - middleware to check if user is logged in
func AuthorizationMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("sessionid")

		if !helpers.ValidateSession(string(cookie.Value)) {
			w.WriteHeader(http.StatusFound)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
