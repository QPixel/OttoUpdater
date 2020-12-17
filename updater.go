package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"io"
	"path/filepath"
	"strings"
	"time"
	"crypto/sha1"
	"bytes"
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
	modpackID		string
)

func init() {
	// Seed random
	rand.Seed(time.Now().Unix())

	// Parse Flags
	flag.StringVar(&modpackID, "modpackid", "", "id of the modpack you are trying to download")
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

	var catalog *Catalog
	var manifest *Manifest

	if manifestID == "" {
		log.Println("Fetching latest Modpack Catalog")

		catalogBytes, err := GetCatalogResponse(modpackID)
		if err != nil {
			log.Fatalf("Failed to fetch catalog: %v", err)
		}
		catalog, err = parseCatalog(catalogBytes)
		if err != nil {
			log.Fatalf("Failed to parse catalog: %v", err)
		}

		if len(catalog.Elements) !=1 || len(catalog.Elements[0].Manifests) < 1 {
			log.Fatal("Unsupported Catalog")
		}

		log.Printf("Catalog %s - %s loaded.\n", catalog.Elements[0].ModpackName, catalog.Elements[0].PackVersion)
	}

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
		
		// log.Println(catalog.Elements[0].ModpackName)
		var manifestBytes []byte
		manifest, manifestBytes, err = fetchManifest(catalog.GetManifestURL())
		if err != nil {
			log.Fatalf("Failed to fetch manifest: %v", err)
		}
		if cachePath != "" {
			ioutil.WriteFile(manifestCachePath, manifestBytes, 0644)
		}
	}

	log.Printf("Manifest %s %s loaded. \n", manifest.ModpackID, manifest.ManifestFileVersion)

	manifestFiles := make(map[string]ManifestFile)
	checkedFiles := make(map[string]ManifestFile)

	for _, file := range manifest.ModFileList {
		manifestFiles[file.FileName] = file
	}

	log.Printf("Found %d files in manifest.\n", len(manifestFiles))

	// for k, file := range manifestFiles {
	// 	func() {
	// 		filePath := filepath.Join(installPath, file.FileName)

	// 		if _, err := os.Stat(filePath); err == nil {
	// 			diskFile, err := os.Open(filePath)
	// 			if err == nil {
	// 				hasher := sha1.New()
	// 				_, err := io.Copy(hasher, diskFile)
	// 				diskFile.Close()

	// 				if err == nil {
	// 					// Compare Checksum
	// 					if bytes.Equal(hasher.Sum(nil), readPackedData(file.FileHash)) {
	// 						log.Printf("File %s found on disk!\n", file.FileName)
	// 						checkedFiles[k] = file
	// 						return
	// 					}
	// 				}
	// 			}
	// 		}

	// 		log.Printf("Downloading File: %s", file.FileName)

	// 		os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	// 	}()
	}
	// log.Println("Found OttoPack v0.9")
	
	// log.Println("Found 100 Mods, 80 Configs")

	// log.Println("Downloading...")

}