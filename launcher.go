package main

import (
	"flag"
	"net/http"	
	"time"
	"math/rand"
	"github.com/Equanox/gotron"
)

var httpClient = &http.Client{}

var (
	manifestID 	string
	installPath 	string
	gamePath 	string
	workerCount	 int
)

const defaultDownloadURL = "https://api.qpixel.me/api/v1/ottomods?version=latest"

func init() {
	// Seed random
	rand.Seed(time.Now().Unix())

	// Parse Flags
	flag.StringVar(&manifestID, "manifest", "", "download a specific manifest")
	// flag.StringVar(&)
	// Set http timeout
	httpClient.Timeout = 30 * time.Second
}


func main() {
	// Make working directories
	// if gamePath != "" {
	// 	os.MkdirAll(gamePath, os.ModePerm)
	// }
	window, err := gotron.New()
	if err != nil {
		panic(err)
	}
	// set default window size
	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 980
	window.WindowOptions.Title = "OttoUpdater"

	// Start the browser window
	done, err := window.Start()
	if err != nil {
		panic(err)
	}

	//Open dev tools after window .start
	window.OpenDevTools()

	<-done
}