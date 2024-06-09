package main

type Product struct {
	name, category string
	price          float64
}

func (p *Product) getName() string {
	return p.name
}

func (p *Product) getCost(_ bool) float64 {
	return p.price
}

// changing (p Product) to (p *Product)
// This is a small change, but it means that the Product type no longer implements
// the Expense interface because the required methods are no longer defined.
// Instead, it is the *Product type that implements the interface, which means that
// pointers to Product values can be treated as Expense values but not regular values.
