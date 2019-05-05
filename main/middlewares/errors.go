package middlewares

import (
	"net/http"
)

//ErrorMiddleware - middlewares for errors
func ErrorMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()

			if err != nil {
				//log.Println(err)
				http.Error(w, "Internal Server Error", 500)
			}

		}()
		handler.ServeHTTP(w, r)
	})
}
