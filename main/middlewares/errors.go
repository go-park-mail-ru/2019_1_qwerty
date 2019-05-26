package middlewares

import (
	"2019_1_qwerty/handlers"
	"net/http"
)

//ErrorMiddleware - middlewares for errors
func ErrorMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()

			if err != nil {
				handlers.ErrorMux(&w, r, http.StatusInternalServerError)
			}

		}()
		handler.ServeHTTP(w, r)
	})
}
