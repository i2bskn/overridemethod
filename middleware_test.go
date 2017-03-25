package overridemethod

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware(t *testing.T) {
	expected := http.MethodPut
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != expected {
			t.Fatal("Middleware: not overwritten")
		}
	})

	mw := Middleware()
	var h http.Handler
	h = mux
	h = mw(h)

	r := httptest.NewRequest(http.MethodPost, "/", nil)
	r.Header.Set(overrideMethodHeader, expected)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
}
