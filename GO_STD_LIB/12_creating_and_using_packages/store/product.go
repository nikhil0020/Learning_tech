package store

// Product describe an item for sale
type Product struct {
	Name, Category string // Name and type of Product
	price          float64
}

// notice that price is in small cap
// means it can only be accessed by method of Product struct
// else if we try to access it we will get error.

// to resolve this issue we need to create a getter for this field.

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (p *Product) Price() float64 {
	return p.price
}

func (p *Product) SetPrice(price float64) {
	p.price = price
}
