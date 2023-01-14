package page

import (
	"bestceramic-parser/models"
	"bestceramic-parser/utils"
	"strconv"

	"github.com/gocolly/colly/v2"
)

func scrapCollections(x *colly.XMLElement, collections []models.Collection) []models.Collection {
	for _, url := range x.ChildAttrs("//div[@class='item__body']//a[@class='item__title' and @href != '/section/']", "href") {
		cInstance := utils.NewCollector()
		collection := Collection(cInstance, utils.Domain+url)
		collections = append(collections, collection)
	}
	return collections
}

func CatalogPlitca(c *colly.Collector, url string, collections []models.Collection) []models.Collection {
	utils.OnRequest(c)
	c.OnXML("//div[@class='contentcols']", func(x *colly.XMLElement) {
		items := x.ChildTexts("//a[@class='pagination__item']")
		if len(items) != 0 {
			lastItem := items[len(items)-1]
			outOfBounds, _ := strconv.Atoi(lastItem)
			for i := 1; i <= outOfBounds; i++ {
				cInstance := utils.NewCollector()
				collections = catalogPage(cInstance, url+"/"+strconv.Itoa(i), collections)
			}
		} else {
			collections = scrapCollections(x, collections)
		}
	})
	c.Visit(url)
	return collections
}

func catalogPage(c *colly.Collector, url string, collections []models.Collection) []models.Collection {
	utils.OnRequest(c)
	c.OnXML("//div[@class='contentcols']", func(x *colly.XMLElement) {
		collections = scrapCollections(x, collections)
	})
	c.Visit(url)
	return collections
}
