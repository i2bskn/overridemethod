package overridemethod

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const unknownMethod = "UNKNOWN"

var methods = []string{
	http.MethodGet,
	http.MethodHead,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
	http.MethodConnect,
	http.MethodOptions,
	http.MethodTrace,
}

var pkgMethods = []int{
	methodGet,
	methodHead,
	methodPost,
	methodPut,
	methodPatch,
	methodDelete,
	methodConnect,
	methodOptions,
	methodTrace,
}

func TestOverrideHTTPRequest(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	r = OverrideHTTPRequest(r)
	if len(Origin(r)) > 0 {
		t.Fatal("overrideRequest: should ignore requests without override method")
	}

	r = httptest.NewRequest(http.MethodPost, "/", nil)
	r.Header.Set(overrideMethodHeader, http.MethodPost)
	r = OverrideHTTPRequest(r)
	if len(Origin(r)) > 0 {
		t.Fatal("overrideRequest: should ignore when no change by overwriting")
	}

	r = httptest.NewRequest(unknownMethod, "/", nil)
	r.Header.Set(overrideMethodHeader, http.MethodPut)
	r = OverrideHTTPRequest(r)
	if len(Origin(r)) > 0 {
		t.Fatal("overrideRequest: should ignore when invalid override method")
	}

	r = httptest.NewRequest(http.MethodPost, "/", nil)
	r.Header.Set(overrideMethodHeader, http.MethodPut)
	r = OverrideHTTPRequest(r)
	if len(Origin(r)) == 0 {
		t.Fatal("overrideRequest: valid override method has not been overwritten")
	}
}

func TestOverrideHTTPMethod(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/", nil)
	expected := ""
	actual := OverrideHTTPMethod(r)
	if expected != actual {
		t.Fatalf("overrideMethod(form none, header none): expected empty, actual %s", actual)
	}

	r = httptest.NewRequest(http.MethodPost, "/", nil)
	expected = http.MethodPut
	r.Header.Set(overrideMethodHeader, strings.ToLower(expected))
	actual = OverrideHTTPMethod(r)
	if expected != actual {
		t.Fatalf("overrideMethod(form none, header %s): expected %s, actual %s", expected, expected, actual)
	}

	r = httptest.NewRequest(http.MethodPost, "/", nil)
	r.ParseForm()
	expected = http.MethodPut
	r.PostForm.Set(overrideMethodParam, strings.ToLower(expected))
	actual = OverrideHTTPMethod(r)
	if expected != actual {
		t.Fatalf("overrideMethod(form %s, header none): expected %s, actual %s", expected, expected, actual)
	}

	r = httptest.NewRequest(http.MethodPost, "/", nil)
	r.ParseForm()
	expected = http.MethodPut
	r.PostForm.Set(overrideMethodParam, expected)
	r.Header.Set(overrideMethodHeader, http.MethodDelete)
	actual = OverrideHTTPMethod(r)
	if expected != actual {
		t.Fatalf("overrideMethod(form %s, header %s): expected %s, actual %s", expected, http.MethodDelete, expected, actual)
	}
}

func TestIsAcceptMethod(t *testing.T) {
	for _, m := range methods {
		if !isAcceptMethod(m) {
			t.Fatalf("should be accept all methods but %s not accept", m)
		}
	}

	if isAcceptMethod(unknownMethod) {
		t.Fatalf("%s is accepted", unknownMethod)
	}
}

func TestParseMethod(t *testing.T) {
	for i := 0; i < len(methods); i++ {
		actual := parseMethod(methods[i])
		if pkgMethods[i] != actual {
			t.Fatalf("parseMethod(%s): expected %v, actual %v", methods[i], pkgMethods[i], actual)
		}
	}

	expected := 0
	actual := parseMethod(unknownMethod)
	if expected != actual {
		t.Fatalf("parseMethod(%s): expected %v, actual %v", unknownMethod, expected, actual)
	}
}
