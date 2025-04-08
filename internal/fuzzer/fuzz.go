package fuzzer

import (
  "fmt"
	"io"
	"net/http"
  "ghozz/internal/tor"
  "ghozz/pkg/output"
  "ghozz/pkg/filter"
	"ghozz/models"
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

func processResponse(resp *http.Response, url string) (models.Result, error) {
	defer resp.Body.Close()
	var resultData models.Result
	
  body, err := io.ReadAll(resp.Body)
  if err != nil {
    return resultData, fmt.Errorf("error reading response: %w", err)
  }

	resultData.URL = url 
	resultData.HttpStatusCode = resp.StatusCode
  
	cfDetected, err := ExtractCloudflareHtml(body)
  if err != nil {
		return resultData, fmt.Errorf("error checking cloudflare: %w", err)
  }
	if cfDetected {
		resultData.CFDetected = true
	} else {
		resultData.CFDetected = false
	}
	
  return resultData, nil
}

func bufferedWriteResults(outputFile string, resultsChan <-chan models.Result, wg *sync.WaitGroup)  {
	defer wg.Done()
 	var results []string
	
 	for result := range resultsChan	 {
 	  results = append(results, result.String())
 	}
 	if err := output.SaveToFile(outputFile, results); err != nil {
 	  fmt.Printf("Error saving output: %v\n", err)
 	}
}

func fuzzWorker(target string, wordlist <-chan string, client *http.Client, hiddenCodes map[int]bool, resultsChan chan<- models.Result, wg *sync.WaitGroup) error {
 	defer wg.Done()

  for word:= range wordlist {
    var url string = fmt.Sprintf("%s/%s", filter.NormalizeURL(target), filter.NormalizeWord(word))
    req, err := http.NewRequest("GET", url,nil)
    if err != nil {
      return fmt.Errorf("Error creating request: %v", err ) 
     }

    resp, err := client.Do(req)
    if err != nil {
      return fmt.Errorf("Error sending request: %v", err)
    }

    if hiddenCodes[resp.StatusCode] {
      resp.Body.Close()
      continue
    }

    result, err := processResponse(resp, word)
    if err == nil {
      fmt.Println(result)
      resultsChan <- result
    }  
  }

	return nil
}
 
func Fuzz(userInput models.CliArgs){ 
	target := userInput.Target
	wordlist := userInput.Wordlist
	httpCode := userInput.HideCode
	clientSetup := userInput.ClientSetup
	torSetup := userInput.TorSetup
	outputFile := userInput.OutputFile
	workers := userInput.Workers

  client := createHTTPClient(clientSetup, torSetup) 
  if torSetup != nil {
    tor.CheckTor(torSetup)
  }
  
  hiddenCodes, err := filter.ParseHideCodes(httpCode)
  if err != nil {
    fmt.Errorf("Error parsing http code: %v", err)
  }
 
  resultsChan := make(chan models.Result, 100)
  wordlistChan := make(chan string, 100)
  var workerWG sync.WaitGroup
	var resultWG sync.WaitGroup

  if outputFile != "" {
    resultWG.Add(1)
    go bufferedWriteResults(outputFile, resultsChan, &resultWG)
  }

  numWorkers := workers
  if workers <= 0 {
    numWorkers = 10
		fmt.Println("[!] Workers must more than 0. Set 10 as default workers...")
  }
  for i:= 0; i < numWorkers; i++ {
  	workerWG.Add(1)
    go fuzzWorker(target, wordlistChan, client, hiddenCodes, resultsChan,  &workerWG)
  }
  for _, word := range wordlist {
    wordlistChan <- word 
  }
  close(wordlistChan)  
 
	workerWG.Wait()
	close(resultsChan)
	 
  if outputFile != "" {
   resultWG.Wait()  
  }
} 
