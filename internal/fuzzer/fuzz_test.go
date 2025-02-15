package fuzzer

import (
  "io"
  "net/http"
  "net/http/httptest"
  "testing"
  "ghozz/internal/tor"
)

func TestFuzz(t *testing.T) {
  // Mock HTTP server
  mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    io.WriteString(w, "Mock response")
  }))
  defer mockServer.Close()

  // Simulate wordlist
  wordlist := []string{"test", "admin", "login"}

  // Use mock HTTP client
  client := mockServer.Client()

  // Call Fuzz with updated parameters (client and no TOR transport)
  Fuzz(mockServer.URL, wordlist, client, nil, "")
}

func TestFuzzTor(t *testing.T) {
  transport, err := tor.SetupTOR()
  if err != nil {
    t.Fatalf("Failed to setup TOR service: %v", err)
  }

  var wordlist = []string{"hidden_service", "secret"}
  
  var target string = "http://check.torproject.org"
  Fuzz(target, wordlist, nil, transport, "")
}
