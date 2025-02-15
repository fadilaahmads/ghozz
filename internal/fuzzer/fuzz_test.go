package fuzzer

import (
  "io"
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestFuzz(t *testing.T) {
  mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    io.WriteString(w, "Mock Response")
  }))

  defer mockServer.Close()

  wordlist := []string{"test", "admin", "login"}

  Fuzz(mockServer.URL, wordlist, nil, "")
}
