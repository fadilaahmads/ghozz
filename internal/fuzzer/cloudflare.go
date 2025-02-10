package fuzzer

import (
  "bytes"
  "fmt"

  "github.com/PuerkitoBio/goquery"
)

func ExtractCloudflareHtml(body []byte) (bool, error) {
  var detected bool
  // detect if the response is actually cloudflare page
  doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
  if err != nil {
    fmt.Errorf("Error parsing HTML: %v", err)
    return false, err
  }

  title := doc.Find("title").Text() 
  if(title == "Attention Required! | Cloudflare"){
    detected = true
    fmt.Println("[!] Cloudflare Detected!")
    fmt.Println("[X] Unable to proceed!")
  } else {
    detected = false
  }
  return detected, nil
 }
