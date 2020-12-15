package main

import (
	"net/http"	
	"time"
	"flag"
	"math/rand"
	"strings"
)

var httpClient = &http.Client{}

const defaultDownloadURL = "https://api.qpixel.me/api/v1/ottomods?version=latest"


var (
	manifestID		string
	downloadURLs	[]string
	version		string
	installPath		string
)

func init() {
	// Seed random
	rand.Seed(time.Now().Unix())

	// Parse Flags
	flag.StringVar(&manifestID, "manifest", "", "download a specific manifest")
	flag.StringVar(&installPath, "install-dir", "files", "folder to write downloaded files to")
	flag.StringVar(&version, "version", "", "downloads a specific version")
	dUrls := flag.String("url", defaultDownloadURL, "downloadUrl")
	flag.Parse()
	downloadURLs = strings.Split(*dUrls, ",")
	// Set http timeout
	httpClient.Timeout = 30 * time.Second
}


func main() {
}