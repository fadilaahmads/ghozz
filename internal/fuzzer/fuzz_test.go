package fuzzer

import ( 
  "net/http"
  "net/http/httptest"
  "strings"
  "testing"
  "ghozz/internal/tor"
)

func TestFuzz(t *testing.T) {
  // Mock HTTP server
  mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path
    switch path {
    case "/login":
      w.WriteHeader(http.StatusMovedPermanently)
    case "/302":
      w.WriteHeader(http.StatusFound)
    case "/404":
      w.WriteHeader(http.StatusNotFound)
    case "/admin":
      w.WriteHeader(http.StatusForbidden)
    case "/405":
      w.WriteHeader(http.StatusMethodNotAllowed)
    case "/503":
      w.WriteHeader(http.StatusServiceUnavailable)
    default:
      w.WriteHeader(http.StatusOK)
    }
  }))
  defer mockServer.Close()
 
  var hideCodes string = "404,403,503"
  t.Log("[*] Hidde codes", strings.Split(hideCodes, ","))

  // Simulate wordlist
  wordlist := []string{"test", "admin", "login","404", "405", "403", "503" } 

  // Use mock HTTP client
  client := mockServer.Client()

  t.Log("Testing Fuzz with no TOR and hide codes") 
  Fuzz(mockServer.URL, wordlist, hideCodes, client, nil, "")
}

func TestFuzzTor(t *testing.T) {
  transport, err := tor.SetupTOR()
  if err != nil {
    t.Fatalf("Failed to setup TOR service: %v", err)
  }

  var wordlist = []string{"hidden_service", "secret"}
  
  var target string = "http://check.torproject.org"
  Fuzz(target, wordlist, "", nil, transport, "")
}
