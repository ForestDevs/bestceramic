package page

import (
	"bestceramic-parser/models"
	"bestceramic-parser/utils"
	"strings"

	"github.com/gocolly/colly/v2"
)

const (
	productMainBlock           = "//div[@class='product-single']"                                                // main block
	productTitle               = "//h1[@itemprop='name']"                                                        // name
	productPrice               = "//div[@class='plate-line__item']"                                              // main price
	productImages              = "//div[@class='product-slider__item'][@data-src]"                               // images                                // product feateruse titiles from product card
	productFeaturesDescription = "//div[@class='communication-prop__col'][1]//p[@class='tile-prop-tabs__value']" // valuse feauters
)

func productFeaturesCollector(x *colly.XMLElement, col map[string]string) map[string]string {
	productFeaturesTitles := x.ChildTexts("//div[@class='product-characteristic__item-text']")
	productFeaturesValues := x.ChildTexts("//div[@class='product-characteristic__item-value']")
	for i := 0; i < len(productFeaturesTitles); i++ {
		col[productFeaturesTitles[i]] = productFeaturesValues[i]
	}
	return col
}

func imagesCollector(x *colly.XMLElement) []string {
	images := make([]string, 0)
	for _, src := range x.ChildAttrs(productImages, "data-src") {
		images = append(images, utils.Domain+src)
	}
	return images
}

func Product(c *colly.Collector, url string) models.Product {
	utils.OnRequest(c)
	var product models.Product
	c.OnXML(productMainBlock, func(x *colly.XMLElement) {
		var productFeatures map[string]string = make(map[string]string)
		name := x.ChildText(productTitle)
		price := strings.ReplaceAll(trimPriceString(x.ChildText(productPrice)), " ", "")
		images := imagesCollector(x)
		productFeatures = productFeaturesCollector(x, productFeatures)
		product = models.NewProduct(name, price, images, productFeatures)
	})
	c.Visit(url)
	return product
}
