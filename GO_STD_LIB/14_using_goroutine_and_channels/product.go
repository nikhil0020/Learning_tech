package main

import "strconv"

type Product struct {
	Name, Category string
	Price          float64
}

var ProductList = []*Product{
	{"Kayak", "Watersports", 279},
	{"Lifejacket", "Watersports", 49.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
}

type ProductGroup []*Product

type ProductData = map[string]ProductGroup

/*

// type ProductData = map[string]ProductGroup

This defines ProductData as an alias for map[string]ProductGroup. It's essentially saying that
ProductData and map[string]ProductGroup are interchangeable. If you use ProductData somewhere, the compiler treats it exactly
 the same as map[string]ProductGroup. There is no new type being created here; it's just another name for the existing type.

// type ProductData map[string]ProductGroup

This defines ProductData as a new type that is based on map[string]ProductGroup. Even though it has the same underlying
structure as map[string]ProductGroup, it is considered a distinct type. This allows you to define methods on ProductData
that wouldn't be applicable to a plain map[string]ProductGroup.

*/

/*

// Alias
type ProductDataAlias = map[string]ProductGroup

// New Type
type ProductDataNewType map[string]ProductGroup

*/

var Products = make(ProductData)

func ToCurrency(val float64) string {
	return "$" + strconv.FormatFloat(val, 'f', 2, 64)
}

func init() {
	for _, p := range ProductList {
		if _, ok := Products[p.Category]; ok {
			Products[p.Category] = append(Products[p.Category], p)
		} else {
			Products[p.Category] = ProductGroup{p}
		}
	}
}
