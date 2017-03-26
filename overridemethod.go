package overridemethod

import (
	"net/http"
	"strings"
)

const (
	overrideMethodParam  = "_method"
	overrideMethodHeader = "X-HTTP-Method-Override"
)

const (
	methodGet = 1 << iota
	methodHead
	methodPost
	methodPut
	methodPatch
	methodDelete
	methodConnect
	methodOptions
	methodTrace
	methodAny = methodGet | methodHead | methodPost | methodPut | methodPatch | methodDelete |
		methodConnect | methodOptions | methodTrace
)

// OverrideHTTPRequest returns the overridden http.Request.
// If not override, returns the received argument as it is.
func OverrideHTTPRequest(r *http.Request) *http.Request {
	if isAcceptMethod(r.Method) {
		method := OverrideHTTPMethod(r)
		if method != r.Method && isAcceptMethod(method) {
			r = setOrigin(r, r.Method)
			r.Method = method
		}
	}

	return r
}

// OverrideHTTPMethod returns the overridden HTTP method.
// If not override, returns the empty string.
func OverrideHTTPMethod(r *http.Request) string {
	if r.Method == http.MethodPost {
		m := r.PostFormValue(overrideMethodParam)
		if len(m) > 0 {
			return strings.ToUpper(m)
		}
	}
	return strings.ToUpper(r.Header.Get(overrideMethodHeader))
}

func isAcceptMethod(m string) bool {
	if (methodAny & parseMethod(m)) == 0 {
		return false
	}
	return true
}

func parseMethod(m string) int {
	switch m {
	case http.MethodGet:
		return methodGet
	case http.MethodHead:
		return methodHead
	case http.MethodPost:
		return methodPost
	case http.MethodPut:
		return methodPut
	case http.MethodPatch:
		return methodPatch
	case http.MethodDelete:
		return methodDelete
	case http.MethodConnect:
		return methodConnect
	case http.MethodOptions:
		return methodOptions
	case http.MethodTrace:
		return methodTrace
	default:
		return 0
	}
}
