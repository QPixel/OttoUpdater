package main

import (
	"encoding/json"
	"os"
	"fmt"
	"io/ioutil"
)

const modpackAPIURl = "https://qpixel.me"

// Catalog defines a catalog 
type Catalog struct {
	Elements []struct {
		ModpackName		string `json:"modpackName"`
		PackVersion		string `json:"packVersion"`
		Manifests		[]struct {
			URI		string `json:"uri"`
		} `json:"manifests"`
	}  `json:"elements"`
	ServerTime	string `json:"serverTime"`
}


// GetCatalogResponse gets the modpack catlog reponse
func GetCatalogResponse(app string) (data[]byte, err error) {
	url := fmt.Sprintf("%s/api/modpacks/%s", modpackAPIURl, app)
	resp, err := httpClient.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Check the reponse code
	if resp.StatusCode != 200 {
		err = fmt.Errorf("invalid status code %d", resp.StatusCode)
		return
	}

	data, err = ioutil.ReadAll(resp.Body)

	return
}

// GetManifestURL returns the Manifest url
func (c *Catalog) GetManifestURL() string {
	for _, m := range c.Elements[0].Manifests {
		return m.URI
	}
	return ""
}

//Load catalog from a file on disk
func readCatalogFile(filename string) (catalog *Catalog, err error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	catalog = new(Catalog)

	err = json.NewDecoder(file).Decode(catalog)
	return
}

func parseCatalog(data []byte) (catalog *Catalog, err error) {
	catalog = new(Catalog)
	// fmt.Println(string(data))
	err = json.Unmarshal(data, catalog)
	return
}