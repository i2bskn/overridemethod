package overridemethod

import (
	"context"
	"net/http"
)

type contextKey int

const (
	originMethodKey contextKey = iota + 1
)

// Origin returns the HTTP method before overriding.
// If not overridden, returns the empty string.
func Origin(r *http.Request) string {
	ctx := r.Context()
	if origin, ok := ctx.Value(originMethodKey).(string); ok {
		return origin
	}
	return ""
}

func setOrigin(r *http.Request, m string) *http.Request {
	ctx := r.Context()
	if _, ok := ctx.Value(originMethodKey).(string); ok {
		return r
	}
	ctx = context.WithValue(ctx, originMethodKey, m)
	return r.WithContext(ctx)
}
