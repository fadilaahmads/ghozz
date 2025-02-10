package output

import (
  "fmt"
  "os"
)

func SaveToFile(filename string, data []string) error{
  file, err := os.Create(filename)
  if err != nil {
    return fmt.Errorf("Error creating file: %v", err)
  }

  defer file.Close()

  for _, line := range data {
    _, err := file.WriteString(line + "\n")
    if err != nil {
      return fmt.Errorf("Error writing to file: %v", err)
    }
  }

  fmt.Printf("[*] Output saved to %s\n", filename)
  return nil
}
