package fuzzer

import (
  "fmt"
	"io"
	"net/http"
  "ghozz/internal/tor"
  "ghozz/pkg/output"
  "ghozz/pkg/filter"
  "time"
  "sync"
)

func createHTTPClient(client *http.Client, torClient *http.Transport) *http.Client {
  if client != nil {
    return client // Use injected client for tests or custom clients.
  }
  return &http.Client{
    Timeout: 15 * time.Second,
    Transport: torClient,
  }
}

func processResponse(resp *http.Response, url string) (string, error) { 
  body, err := io.ReadAll(resp.Body)
  if err != nil {
    return "", fmt.Errorf("error reading response: %w", err)
  }

  cfDetected, _ := ExtractCloudflareHtml(body)
  if cfDetected {
    return "[!] Cloudflare detected, skipping", nil
  }
  resp.Body.Close() 

  return fmt.Sprintf("[*] URL: %s | Status: %d", url, resp.StatusCode), nil
}

func bufferedWriteResults(outputFile string, resultsChan <-chan string, wg *sync.WaitGroup)  {
 defer wg.Done()
 var results []string
 for result := range resultsChan {
   results = append(results, result)
 }
 if err := output.SaveToFile(outputFile, results); err != nil {
   fmt.Printf("Error saving output: %v\n", err)
 }
}

func fuzzWorker(target string, wordlist <-chan string, client *http.Client, hiddenCodes map[int]bool, resultsChan chan<- string, wg *sync.WaitGroup) {
  defer wg.Done()
  for word:= range wordlist {
    var url string = fmt.Sprintf("%s/%s", target, word) 
    req, err := http.NewRequest("GET", url,nil)
    if err != nil {
      fmt.Println("Error creating request: ", err )
      return
     }

    resp, err := client.Do(req)
    if err != nil {
      fmt.Println("Error sending request: ", err)
      return
    }

    if hiddenCodes[resp.StatusCode] {
      resp.Body.Close()
      continue
    }

    result, err := processResponse(resp, url)
    if err == nil {
      fmt.Println(result)
      resultsChan <- result
    }  
  }
}

func Fuzz(target string, wordlist []string, httpCode string, clientSetup *http.Client, torSetup *http.Transport, outputFile string, workers int) { 
  client := createHTTPClient(clientSetup, torSetup) 
  if torSetup != nil {
    tor.CheckTor(torSetup)
  }
  
  hiddenCodes, err := filter.ParseHideCodes(httpCode)
  if err != nil {
    fmt.Errorf("Error parsing http code: %v", err)
  }
 
  resultsChan := make(chan string, 100)
  wordlistChan := make(chan string, 100)
  var wg sync.WaitGroup

  if outputFile != "" {
    wg.Add(1)
    go bufferedWriteResults(outputFile, resultsChan, &wg)
  }

  numWorkers := 10
  if workers <= 0 {
    numWorkers = workers
  }
  for i:= 0; i < numWorkers; i++ {
    wg.Add(1)
    go fuzzWorker(target, wordlistChan, client, hiddenCodes, resultsChan,  &wg)
  }
  for _, word := range wordlist {
    wordlistChan <- word 
  }
  close(wordlistChan)  
 
  wg.Wait()

  if outputFile != "" {
    close(resultsChan) 
  }
} 
