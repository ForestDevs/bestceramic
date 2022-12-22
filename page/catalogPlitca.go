package page

import (
	"bestceramic-parser/models"
	"bestceramic-parser/utils"

	"github.com/gocolly/colly/v2"
)

func scrapCollections(x *colly.XMLElement, collections []models.Collection) []models.Collection {
	for _, url := range x.ChildAttrs("//div[@class='item__body']//a[@class='item__title']", "href") {
		cInstance := utils.NewCollector()
		collection := Collection(cInstance, utils.Domain+url)
		collections = append(collections, collection)
	}
	return collections
}

func CatalogPlitca(c *colly.Collector, url string) []models.Collection {
	var collections []models.Collection = make([]models.Collection, 0)
	utils.OnRequest(c)
	c.OnXML("//div[@class='contentcols']", func(x *colly.XMLElement) {
		nextPage := x.ChildAttr("//a[@class='pagination__item pagination__item_next']", "href")
		collections = scrapCollections(x, collections)
		if nextPage != "" {
			CatalogPlitca(c, utils.Domain+nextPage)
		}
	})
	c.Visit(url)
	return collections
}
