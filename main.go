package main

import (
  "bufio"
  "fmt"
  "context"
  "bytes"
  "net"
  "net/http"
  "net/url"
  "golang.org/x/net/proxy"
  "io"
  "os"
  "time"
  "github.com/PuerkitoBio/goquery"
  "github.com/urfave/cli/v2"
)
  

func CheckTor(tor *http.Transport){
  client := &http.Client{
    Transport: tor,
  }
  // Make a request
	resp, err := client.Get("https://check.torproject.org")  
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

  ExtractTorHtml(body)
}

func SetupTOR() (*http.Transport, error) {
  torProxy := "socks5://127.0.0.1:9050"
  proxyURL, err := url.Parse(torProxy)
  if err != nil {
    errorMessage := fmt.Errorf("Error parsing proxy URL: %v", err) 
    return nil,errorMessage
  }

  dialer, err := proxy.FromURL(proxyURL, proxy.Direct)
  if err != nil {
    return nil, fmt.Errorf("error creating proxy dialer: %v", err)
  } 
  
  dialContext := func(ctx context.Context, network, addr string)(net.Conn, error){
    return dialer.Dial(network, addr)
  }

  transport := &http.Transport{
    DialContext:  dialContext,
    ForceAttemptHTTP2:  true,
    MaxIdleConns: 10,
    IdleConnTimeout:  30 * time.Second,
    TLSHandshakeTimeout:  10 * time.Second,
    ExpectContinueTimeout:  1 * time.Second,
  }

  return transport,nil
}

func ExtractTorHtml(body []byte){
  doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
  if err != nil {
    fmt.Errorf("Error parsing HTML: %v", err)
    return
  }

  title := doc.Find("title").Text()
  ip := doc.Find("strong").Text()
  fmt.Println("[*] Status: ", title)
  fmt.Println("[*] IP: ", ip)
}

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

func Fuzz(target string, wordlist []string, tor *http.Transport) {
  const constUrl = "%s/%s"
  var client *http.Client = &http.Client{
    Transport: tor,
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
    // body, err := io.ReadAll(resp.Body)
    // if err != nil {
    //   fmt.Println("[X]Error reading response: ", err)
    //   return
    // }
    // fmt.Println("[*]Response: ", string(body))
  }
}

func main()  {
  fmt.Println("GoFuzzer")
  tor, err := SetupTOR()
  if err != nil {
    fmt.Errorf("Error setting up TOR proxy: %v", err)
  }
  
  CheckTor(tor)

  app := &cli.App{
    Name: "GoFuzzer",
    Usage: "Directory Fuzzing With TOR Support",
    Flags: []cli.Flag{
      &cli.StringFlag{Name: "target", Required: true},
      &cli.StringFlag{Name: "wordlist", Required: true},
    },
    Action: func(c *cli.Context) error {
      target:= c.String("target")
      wordlistFile := c.String("wordlist")
      wordlist, err := ReadWordlist(wordlistFile)
      if err != nil {
        return err
      }
      Fuzz(target, wordlist, tor)
      return nil
    },
  }

  err = app.Run(os.Args)
  if err != nil {
    fmt.Errorf("Error running app: %v", err)
  }
}
