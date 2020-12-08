package main

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)

// ManifestFile defines a a file within FileManifestList
type ManifestFile struct {
	Modname	string	`json:"Modname"`
	FileName 	string `json:"Filename"`
	FileVersion	string	`json:"FileVersion"`
	FileID	int	`json:"FileID"`
	FileURL	string `json:"FileUrl"`
	FileHash string `json:"FileHash"`
}

// Manifest defines a manifest or whatever
type Manifest struct {
	ManifestFileVersion string 	`json:"ManifestFileVersion"`
	ModpackID	string 	`json:"ModpackID"`
	LaunchCommand	string	`json:"LaunchCommand"`
	FileManifestList []ManifestFile	`json:"FileManifestList"`
}

func readManifestFile(filename string) (*Manifest , error) {
	//Open File
	file, err := os.Open(filename)
	if err !=nil {
		return nil, err
	}
	defer file.Close()

	manifest:= new(Manifest)

	if err := json.NewDecoder(file).Decode(manifest); err != nil {
		return nil, err
	}

	return manifest, nil
}

func fetchManifest(url string) (manifest *Manifest, body []byte, err error) {
	// Get manifest
	resp, err := httpClient.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Check response code
	if resp.StatusCode != 200 {
		err = fmt.Errorf("invalid status code %d", resp.StatusCode)
		return
	}

	// Read body
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	
	// Create a new manifest instance
	manifest = new(Manifest)

	// Parse response body
	err = json.Unmarshal(body, manifest)
	return
}