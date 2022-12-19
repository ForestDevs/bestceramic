package main

import (
	"bestceramic-parser/page"
	"bestceramic-parser/utils"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Collections []string `json: "collections"`
	Brands      []string `json: "brands"`
}

func main() {

	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	cfg := Config{}

	if err := json.Unmarshal(file, &cfg); err != nil {
		panic(err)
	}

	// s := spinner.New(spinner.CharSets[80], 100*time.Millisecond) // Build our new spinner
	// s.Start()
	c := utils.NewCollector()

	if len(cfg.Brands) != 0 {
		for _, url := range cfg.Brands {
			page.CatalogPlitca(c, url)
		}

	} else if len(cfg.Collections) != 0 {
		for _, url := range cfg.Collections {
			page.Collection(c, url)
		}
	}

	// s.Stop()
}
