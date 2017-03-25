package overridemethod

import (
	"net/http"
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

func overrideRequest(r *http.Request) *http.Request {
	if isAcceptMethod(r.Method) {
		method := overrideMethod(r)
		if method != r.Method && isAcceptMethod(method) {
			r = setOrigin(r, r.Method)
			r.Method = method
		}
	}

	return r
}

func overrideMethod(r *http.Request) (m string) {
	m = r.FormValue(overrideMethodParam)
	if len(m) > 0 {
		return
	}
	m = r.Header.Get(overrideMethodHeader)
	return
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
