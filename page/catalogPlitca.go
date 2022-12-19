package page

import (
	"bestceramic-parser/utils"

	"github.com/gocolly/colly/v2"
)

func scrapCollections(x *colly.XMLElement) {
	for _, url := range x.ChildAttrs("//div[@class='item__body']//a[@class='item__title']", "href") {
		cInstance := utils.NewCollector()
		Collection(cInstance, utils.Domain+url)
	}
}

func CatalogPlitca(c *colly.Collector, url string) {
	utils.OnRequest(c)
	c.OnXML("//div[@class='contentcols']", func(x *colly.XMLElement) {
		nextPage := x.ChildAttr("//a[@class='pagination__item pagination__item_next']", "href")
		scrapCollections(x)
		if nextPage != "" {
			CatalogPlitca(c, utils.Domain+nextPage)
		}
	})
	c.Visit(url)
}
