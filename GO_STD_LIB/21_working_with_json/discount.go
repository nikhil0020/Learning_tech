package main

type DiscountedProduct struct {
	*Product
	Discount float64
}

type DiscountedProductCustomField struct {
	*Product `json:"product"`
	Discount float64
}

type DiscountedProductOmitField struct {
	*Product `json:"product"`
	Discount float64 `json:"-"`
}

type DiscountedProductOmitUnassigned struct {
	*Product `json:"product,omitempty"`
	Discount float64 `json:",string"`
}
