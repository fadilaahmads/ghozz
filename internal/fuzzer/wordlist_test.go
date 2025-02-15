package fuzzer

import (
  "fmt"
  "testing"
  "os"
)

func TestReadWordlist(t *testing.T) {
  file, err := os.CreateTemp("", "wordlist")
  if err != nil {
    fmt.Println("Error creating file: ",err)
    t.Fatal(err)
  }

  defer os.Remove(file.Name())

  _, _ = file.WriteString("admin\nlogin\nsecret\n")
  file.Close()

  words, err := ReadWordlist(file.Name())
  if err != nil {
    t.Errorf("Unexpected error: %v", err)
  }

  expected := []string{"admin", "login", "secret"}
  for i, word := range words {
    if word != expected[i] {
      t.Errorf("Expected %s, got %s", expected[i], word)
    }
  }
} 
