package fuzzer

import (
  "bytes"
  "fmt"

  "github.com/PuerkitoBio/goquery"
)

func ExtractCloudflareHtml(body []byte) (bool, error) { 
  // detect if the response is actually cloudflare page
  doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
  if err != nil {
    return false, fmt.Errorf("Error parsing HTML: %v", err) 
  }

  title := doc.Find("title").Text() 
  if(title == "Attention Required! | Cloudflare"){
		return true, nil
  }

  return false, nil
 }
