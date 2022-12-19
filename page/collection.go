package page

import (
	"bestceramic-parser/models"
	"bestceramic-parser/utils"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

const (
	collectionMainBlock    = "//div[@class='product-page']"                                     // collection cards block
	collectionTitle        = "//h1"                                                             // collection name
	collectionPrice        = "//div[@class='product-info__price']"                              // price
	collectionImage        = "//div[@class='product-slider__item' and contains(@data-src,'/')]" //images from slide
	barnd                  = "//span[@itemprop='brand']"                                        // brand                                                                                                            // first collection image
	collectionProductsCard = "//div[@class='item__body']//a[@class='item__title']"              // products cards hrefs
)

func productsCollector(x *colly.XMLElement) []models.Product {
	var products []models.Product
	for _, cardHref := range x.ChildAttrs(collectionProductsCard, "href") {
		cInstance := utils.NewCollector()
		products = append(products, Product(cInstance, utils.Domain+cardHref))
	}
	return products
}

func collectionImageCollector(x *colly.XMLElement) string {
	var images []string
	for _, image := range x.ChildAttrs(collectionImage, "data-src") {
		images = append(images, utils.Domain+image)
	}
	return strings.Join(images, ";")
}

func trimPriceString(str string) string {
	r := regexp.MustCompile("\n\\s+")
	return strings.ReplaceAll(r.ReplaceAllString(str, ""), "Цена от: ", "")
}

func Collection(c *colly.Collector, url string) {
	utils.OnRequest(c)

	c.OnXML(collectionMainBlock, func(x *colly.XMLElement) {
		name := x.ChildText(collectionTitle)
		price := trimPriceString(x.ChildText(collectionPrice))
		image := collectionImageCollector(x)
		brand := x.ChildText(barnd)
		products := productsCollector(x)
		collection := models.NewCollection(name, price, image, brand, products)
		utils.ExcelWrite(collection)
	})

	c.Visit(url)
}
