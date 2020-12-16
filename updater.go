package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var httpClient = &http.Client{}


const defaultDownloadURL = "https://media.forgecdn.net/"
const configDownloadURL = "https://cdn.qpixel.me/otto/configs"
const isDev = true;
var (
	manifestID		string
	downloadURLs	[]string
	version		string
	installPath		string
	cachePath 	string
)

func init() {
	// Seed random
	rand.Seed(time.Now().Unix())

	// Parse Flags
	flag.StringVar(&manifestID, "manifest", "", "download a specific manifest")
	flag.StringVar(&installPath, "install-dir", "files", "folder to write downloaded files to")
	flag.StringVar(&version, "version", "", "downloads a specific version")
	flag.StringVar(&cachePath, "cachepath", "", "place to store cache")
	dUrls := flag.String("url", defaultDownloadURL, "downloadUrl")
	flag.Parse()
	downloadURLs = strings.Split(*dUrls, ",")
	// Set http timeout
	httpClient.Timeout = 30 * time.Second
}


func main() {

	if cachePath != "" {
		os.MkdirAll(cachePath, os.ModePerm)
	}


	var manifest *Manifest

	manifestCachePath := filepath.Join(cachePath, "manifest.json")
	if manifestID != "" {
		log.Fatal("Fetching a Specific Manifest is not implimented yet.")

		// var err error
		// manifest, _, err = fetchManifest(fmt.Sprintf("%s/"))
	} else if _, err := os.Stat(manifestCachePath); err == nil && cachePath != "" {
		log.Printf("Loading manifest from cache...")

		manifest, err = readManifestFile(manifestCachePath)
		if err != nil {
			log.Fatalf("Failed to read manifest %v", err)
		}
	} else {
		log.Println("Loading latest Manifest")
		
		var manifestBytes []byte
		manifest, manifestBytes, err = fetchManifest(defaultDownloadURL)
		if err != nil {
			log.Fatalf("Failed to fetch manifest: %v", err)
		}
		if cachePath != "" {
			ioutil.WriteFile(manifestCachePath, manifestBytes, 0644)
		}
	}

	log.Printf("Manifest %s %s loaded. \n", manifest.ModpackID, manifest.ManifestFileVersion)

	manifestFiles := make(map[string]ManifestFile)
	// checkedFiles := make(map[string]ManifestFile)

	for _, file := range manifest.FileManifestList {
		manifestFiles[file.FileName] = file
	}

	log.Printf("Found %d files in manifest.\n", len(manifestFiles))
	// log.Println("Found OttoPack v0.9")
	
	// log.Println("Found 100 Mods, 80 Configs")

	// log.Println("Downloading...")

}