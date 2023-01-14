package page

import (
	"bestceramic-parser/models"
	"bestceramic-parser/utils"
	"fmt"
	"strconv"

	"github.com/gocolly/colly/v2"
)

func scrapCollections(x *colly.XMLElement, collections []models.Collection) []models.Collection {
	brandName := x.ChildAttrs("/../../../..//ul[contains(@class,'breadcrumbs')]/li", "data-text")
	fmt.Println(brandName[len(brandName)-1])
	findOpts := fmt.Sprintf("//div[@class='item__body']//a[@class='item__title' and @href != '/section/']/../dl//a[contains( @title, '%s')]/../../../a[@class='item__title']", brandName[len(brandName)-1])
	for _, url := range x.ChildAttrs(findOpts, "href") {
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
