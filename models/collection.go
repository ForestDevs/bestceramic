package models

type Collection struct {
	Name     string
	Price    string
	Image    string
	Brand    string
	Products []Product
}

func NewCollection(name string, price string, image string, brand string, products []Product) Collection {
	return Collection{
		Name:     name,
		Price:    price,
		Image:    image,
		Brand:    brand,
		Products: products,
	}
}
