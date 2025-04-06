package filter

import (
  "fmt"
  "strings"
)

func NormalizeURL(url string) string {
	return strings.TrimRight(url, "/")
}

func NormalizeWord(word string) string {
	return strings.TrimLeft(word, "/")
}

func ParseHideCodes(flag string) (map[int]bool, error) {
  hiddenCodes := make(map[int]bool)
  invalidCodes := []string{}
  
  if flag == "" {
    return nil, fmt.Errorf("Empty codes")

  }

  for _, code := range strings.Split(flag, ",") {
    var c int
    _, err := fmt.Sscanf(code, "%d", &c)
    if err != nil {
      invalidCodes = append(invalidCodes, code)
    } else {
      hiddenCodes[c] = true
    }
  }

  if len(invalidCodes) > 0 {
    return hiddenCodes, fmt.Errorf("invalid status codes provided: %v", strings.Join(invalidCodes, ", "))
  }
   
  return hiddenCodes, nil
}
