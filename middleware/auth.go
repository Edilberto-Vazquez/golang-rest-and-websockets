package middleware

import (
	"net/http"

	"github.com/Edilberto-Vazquez/golang-rest-and-websockets/server"
	"github.com/Edilberto-Vazquez/golang-rest-and-websockets/utils"
)

var (
	NO_AUTH_NEEDED = []string{
		"login",
		"signup",
	}
)

func CheckAuthMiddleware(s server.Server) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !utils.RouteNeedToken(r.URL.Path, NO_AUTH_NEEDED) {
				next.ServeHTTP(w, r)
				return
			}
			_, err := utils.ProcessToken(r.Header.Get("Authorization"), s)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
