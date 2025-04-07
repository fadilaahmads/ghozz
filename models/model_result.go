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

func (r models.Result) ShowResult() string {
	fm.Println("[>] URL: %s", r.URL)
	fmt.Println("[>] Code: %d", r.HttpStatusCode)
	if r.CFDetected {
		fmt.Println("[!] Cloudflare detected! skipping...")
	}
}
