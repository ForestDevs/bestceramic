package main

import (
	"bestceramic-parser/models"
	"bestceramic-parser/page"
	"bestceramic-parser/utils"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Collections []string `json:"collections"`
	Brands      []string `json:"brands"`
	UpdateKeys  bool     `json:"updateKeys"`
}

var collectionsKeys map[string]int
var productsKeys map[string]int

func main() {

	//cfg
	file, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	cfg := Config{}

	if err := json.Unmarshal(file, &cfg); err != nil {
		panic(err)
	}

	//collections
	fileCollections, err := os.ReadFile("./collectionsKeys.json")
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(fileCollections, &collectionsKeys); err != nil {
		panic(err)
	}

	//collections
	fileProducts, err := ioutil.ReadFile("./productsKeys.json")
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(fileProducts, &productsKeys); err != nil {
		panic(err)
	}

	// s := spinner.New(spinner.CharSets[80], 100*time.Millisecond) // Build our new spinner
	// s.Start()
	c := utils.NewCollector()

	if len(cfg.Brands) != 0 {
		for _, url := range cfg.Brands {
			utils.ExcelWriteMultipleData(page.CatalogPlitca(c, url), collectionsKeys, productsKeys, cfg.UpdateKeys)
		}

	} else if len(cfg.Collections) != 0 {
		for _, url := range cfg.Collections {
			var collections []models.Collection = make([]models.Collection, 1)
			collections = append(collections, page.Collection(c, url))
			utils.ExcelWriteMultipleData(collections, collectionsKeys, productsKeys, cfg.UpdateKeys)
		}
	}

	// s.Stop()
}
