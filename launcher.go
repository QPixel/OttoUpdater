package main

import (
	"net/http"	
	"time"
	"math/rand"
)

var httpClient = &http.Client{}

const defaultDownloadURL = "https://api.qpixel.me/api/v1/ottomods?version=latest"

func init() {
	// Seed random
	rand.Seed(time.Now().Unix())

	// Parse Flags
	// flag.StringVar(&manifestID, "manifest", "", "download a specific manifest")
	// flag.StringVar(&)
	// Set http timeout
	httpClient.Timeout = 30 * time.Second
}


func main() {
}