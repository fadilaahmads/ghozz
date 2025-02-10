package fuzzer

import (
  "fmt"
	"io"
	"net/http"
  "ghozz/internal/tor"
  "ghozz/pkg/output"
)

func Fuzz(target string, wordlist []string, torSetup *http.Transport, outputFile string) {
  tor.CheckTor(torSetup)
  
  var results []string

  const constUrl = "%s/%s"
  var client *http.Client = &http.Client{
    Transport: torSetup,
  }
  for _, word:= range wordlist {
    var url string = fmt.Sprintf(constUrl, target, word) 
    req, err := http.NewRequest("GET", url,nil)
    if err != nil {
      fmt.Println("Error creating request: ", err )
     return
     }

    //req.Header.Set("Content-Type", "application/json")

    // Use hte HTTP client  to send the request 
    resp, err := client.Do(req)
    if err != nil {
     fmt.Println("Error sending request: ", err)
      return
    }

    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
      fmt.Println("[X]Error reading response: ", err)
      return
    }
    CFDetected, err := ExtractCloudflareHtml(body)
    if err != nil {
      fmt.Println("Error extracting cloudflare page: ", err)
    }
    if CFDetected == true {
      continue
    }
    result := fmt.Sprintf("[*] URL: %s | Status: %d", url, resp.StatusCode)
    fmt.Println(result)
    results = append(results, result)

    // print blank line
    fmt.Println("") 
  } 

  if outputFile != "" {
    err := output.SaveToFile(outputFile, results)
      if err != nil {
        fmt.Println("Error saving output: ", err)
      }
    }
}
