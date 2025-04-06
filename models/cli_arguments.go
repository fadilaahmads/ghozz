package models

import (
	"net/http"
)

type CliArgs struct {
	Target string 
	Wordlist []string 
	OutputFile string
	ClientSetup *http.Client 
	TorSetup *http.Transport
	HideCode string
	Workers int
} 
