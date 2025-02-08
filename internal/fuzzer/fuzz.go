package fuzzer

import (
  "fmt"
	"io"
	"net/http"
  "gofuzzer/internal/tor"
)

func Fuzz(target string, wordlist []string, torSetup *http.Transport) {
  tor.CheckTor(torSetup)

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

    fmt.Println("[*]URL: ", url)
    fmt.Println("[*]Status Code: ", resp.StatusCode)
    // fmt.Println("[|]Payload: ", string(data))
    body, err := io.ReadAll(resp.Body)
    if err != nil {
      fmt.Println("[X]Error reading response: ", err)
      return
    }
    // fmt.Println("[*]Response: ", string(body))
    ExtractCloudflareHtml(body)
    // print blank line
    fmt.Println("")
  }
}
