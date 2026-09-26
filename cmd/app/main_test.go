package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func code(t *testing.T, path string) int {
	t.Helper()
	resp, err := newApp().Test(httptest.NewRequest(http.MethodGet, path, nil))
	if err != nil {
		t.Fatal(err)
	}
	return resp.StatusCode
}

func TestServesAtRoot(t *testing.T) {
	for _, path := range []string{"/health", "/"} {
		if got := code(t, path); got != http.StatusOK {
			t.Fatalf("GET %s = %d", path, got)
		}
	}
}
