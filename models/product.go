package models

type Product struct {
	Name     string
	Price    string
	Images   []string
	Features map[string]string
}

func NewProduct(name string, price string, images []string, features map[string]string) Product {
	return Product{
		Name:     name,
		Price:    price,
		Images:   images,
		Features: features,
	}
}
