package overridemethod

import (
	"net/http"
)

func Middleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, overrideRequest(r))
		})
	}
}
