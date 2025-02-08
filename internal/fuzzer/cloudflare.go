package fuzzer

import (
  "bytes"
  "fmt"

  "github.com/PuerkitoBio/goquery"
)

func ExtractCloudflareHtml(body []byte){
  // detect if the response is actually cloudflare page
  doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
  if err != nil {
    fmt.Errorf("Error parsing HTML: %v", err)
    return
  }

  title := doc.Find("title").Text() 
  if(title == "Attention Required! | Cloudflare"){
    fmt.Println("[!] Cloudflare Detected!")
    fmt.Println("[X] Unable to proceed!")
  }
 }
