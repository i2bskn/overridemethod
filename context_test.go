package overridemethod

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOrigin(t *testing.T) {
	expected := ""
	r := httptest.NewRequest(http.MethodPut, "/", nil)
	actual := Origin(r)
	if expected != actual {
		t.Fatalf("unexpected original HTTP method: expected %s, actual %s", expected, actual)
	}

	expected = http.MethodPost
	r = setOrigin(r, expected)
	actual = Origin(r)
	if expected != actual {
		t.Fatalf("unexpected original HTTP method: expected %s, actual %s", expected, actual)
	}

	r = setOrigin(r, http.MethodDelete)
	actual = Origin(r)
	if expected != actual {
		t.Fatalf("unexpected original HTTP method: expected %s, actual %s", expected, actual)
	}
}
