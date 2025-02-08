package fuzzer

import (
  "bufio"
  "fmt"
  "os"
)

func ReadWordlist(filePath string) ([]string, error) {
  file ,err := os.Open(filePath)
  if err != nil {
    fmt.Errorf("Error opening file: %v", err)
  }

  defer file.Close()

  var words []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan(){
    words = append(words, scanner.Text())
  }

  return words, scanner.Err()
}
