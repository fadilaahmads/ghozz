package output

import (
  "bufio"
  "os"
  "testing"
)

func TestSaveToFile(t *testing.T)  {
  var filename string = "temp.txt"
  var data = []string{"output1", "output2", "output3"}

  err := SaveToFile(filename, data)
  if err != nil {
    t.Errorf("Failed to save output to file: %v", err)
  }

  expected := []string{"output1", "output2", "output3"} 

  file, err := os.Open(filename)
  if err != nil {
    t.Errorf("Error opening file: %v", err)
  }
  defer file.Close()
  defer os.Remove(filename)

  var words []string

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)
  for scanner.Scan() {
    words = append(words, scanner.Text())
  } 
  
  for i,word := range words {
    if word != expected[i]{
      t.Errorf("Expected %s, got %s", expected, word)
    }
  }
}
