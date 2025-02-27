package fuzzer

import (
  "fmt"
	"io"
	"net/http"
  "ghozz/internal/tor"
  "ghozz/pkg/output"
  "ghozz/pkg/filter"
  "time"
)

func getClient(client *http.Client, torClient *http.Transport) *http.Client {
  if client != nil {
    return client // Use injected client for tests or custom clients.
  }
  return &http.Client{
    Timeout: 15 * time.Second,
    Transport: torClient,
  }
}

func Fuzz(target string, wordlist []string, httpCode string, clientSetup *http.Client, torSetup *http.Transport, outputFile string) { 
  client := getClient(clientSetup, torSetup)
  
  if torSetup != nil {
    tor.CheckTor(torSetup)
  }
   
  results := make([]string, 0, len(wordlist)) 

  const constUrl = "%s/%s"

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
    
    hiddenCodes, err := filter.ParseHideCodes(httpCode)
    if err != nil {
      fmt.Errorf("Error parsing http code: %v", err)
    }
    if hiddenCodes[resp.StatusCode] {
      continue
    }

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
    
   resp.Body.Close()

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
