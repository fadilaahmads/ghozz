package tor

import (
  "bytes"
  "io"
  "fmt"
  "net/http"

  "github.com/PuerkitoBio/goquery"
)

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
