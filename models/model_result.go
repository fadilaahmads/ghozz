package models

import (
	"fmt"
)

type Result struct {
	URL string
	HttpStatusCode int
	CFDetected bool
}

func (r Result) String() string {
	if r.CFDetected {
		return fmt.Sprintf("[!] URL: %s | Code: %d | CLOUDFLARE DETECTED!", r.URL, r.HttpStatusCode)
	} else {
		return fmt.Sprintf("[*] URL: %s | Code: %d", r.URL, r.HttpStatusCode)
	}
}
