package main

import (
	"net/http"
	"net/http/httptest"
	"os"
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

func TestServesAtRootWhenUnset(t *testing.T) {
	os.Unsetenv("BASE_PATH")
	if got := code(t, "/health"); got != http.StatusOK {
		t.Fatalf("GET /health = %d", got)
	}
}

func TestServesUnderPrefix(t *testing.T) {
	t.Setenv("BASE_PATH", "/direct/agent-7:3000")
	if got := code(t, "/direct/agent-7:3000/health"); got != http.StatusOK {
		t.Fatalf("prefixed = %d", got)
	}
	if got := code(t, "/health"); got != http.StatusNotFound {
		t.Fatalf("bare = %d, want 404", got)
	}
}
