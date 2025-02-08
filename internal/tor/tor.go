package tor 

import (
  "context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/proxy"
)

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
