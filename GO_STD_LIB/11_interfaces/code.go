package main

import "fmt"

type ProductList []Product

func (products *ProductList) calcCategoryTotal() map[string]float64 {
	totals := make(map[string]float64)

	for _, p := range *products {
		totals[p.category] += p.price
	}

	return totals
}

// The type keyword is used to create an alias for the []Product type, with the name ProductList.
// This type can be used to define methods, either directly for value type receivers or with a pointer, as in this example.

func DefiningMethodForTypeAliases() {
	products := ProductList{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}

	for category, total := range products.calcCategoryTotal() {
		fmt.Println("Category: ", category, "Total: ", total)
	}
}
