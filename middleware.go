package overridemethod

import (
	"net/http"
)

// Middleware returns function to wrap with HTTP method overriding.
// Wrapping http.Handler rewrites the HTTP method before http.Handler's ServeHTTP call.
//	mw := overridemethod.Middleware()
//	var handler http.Handler
//	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
//	handler = mw(handler)
//	handler.ServeHTTP(w, r)
func Middleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, OverrideHTTPRequest(r))
		})
	}
}
