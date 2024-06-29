package main

import "fmt"

func calcWithTax(price float64) float64 {
	return price + (price * 0.2)
}
func calcWithoutTax(price float64) float64 {
	return price
}

func Normal() {
	fmt.Println("Function with no parameter and no returns")
}

func UnderstandingFunctionTypes() {

	var nor func() = Normal
	nor()

	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		var calcFunc func(float64) float64
		if price > 100 {
			calcFunc = calcWithTax
		} else {
			calcFunc = calcWithoutTax
		}
		totalPrice := calcFunc(price)
		fmt.Println("Product:", product, "Price:", totalPrice)
	}
}

func UnderstandingFunctionComparision() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		var calcFunc func(float64) float64
		fmt.Println("Function assigned:", calcFunc != nil)
		if price > 100 {
			calcFunc = calcWithTax
		} else {
			calcFunc = calcWithoutTax
		}
		fmt.Println("Function assigned:", calcFunc != nil)
		totalPrice := calcFunc(price)
		fmt.Println("Product:", product, "Price:", totalPrice)
	}
}

// Using Functions as Arguments

func printPrice(product string, price float64, calculator func(float64) float64) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

func UsingFunctionsAsArguments() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		if price > 100 {
			printPrice(product, price, calcWithTax)
		} else {
			printPrice(product, price, calcWithoutTax)
		}
	}
}

// Using Function as Result

func selectCalculator(price float64) func(float64) float64 {
	if price > 100 {
		return calcWithTax
	}
	return calcWithoutTax
}

func UsingFunctionAsResult() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		printPrice(product, price, selectCalculator(price))
	}
}

// Creating Function Type Aliases

type calcFunc func(float64) float64

func printPrice2(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}
func selectCalculator2(price float64) calcFunc {
	if price > 100 {
		return calcWithTax
	}
	return calcWithoutTax
}

func CreatingFunctionTypeAliases() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		printPrice2(product, price, selectCalculator2(price))
	}
}

// Using the Functional Literals

func selectCalculatorUsingFunctionLiterals(price float64) calcFunc {
	if price > 100 {
		var withTax calcFunc = func(price float64) float64 {
			return price + (price * 0.2)
		}
		return withTax
	}
	withoutTax := func(price float64) float64 {
		return price
	}
	return withoutTax
}

func UsingFunctionalLiterals() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		printPrice(product, price, selectCalculatorUsingFunctionLiterals(price))
	}
}

func main() {
	// UnderstandingFunctionTypes()
	// UnderstandingFunctionComparision()
	// UsingFunctionsAsArguments()
	// CreatingFunctionTypeAliases()
	UsingFunctionalLiterals()
}
