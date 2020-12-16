package main

type Catalog struct {
	Elements []struct {
		ModpackName		string `json:"modpackName"`
		LabelName		string `json:"labelName"`
		PackVersion		string `json:"buildVersion"`
		Manifests		[]struct {
			URI		string `json:"uri"`
		} `json:"manifests"`
	}  `json:"elements"`
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
	file
}