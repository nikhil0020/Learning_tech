package main

import (
	"fmt"
	"time"
)

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	// var resultChannel chan float64 = make(chan float64)
	var resultChannel chan float64 = make(chan float64, 2)

	for category, group := range data {
		go group.TotalPrice(category, resultChannel)
	}

	for i := 0; i < len(data); i++ {
		fmt.Println("-- channel read pending", len(resultChannel), "items in buffer, size", cap(resultChannel))
		value := <-resultChannel
		fmt.Println("-- channel read complete", value)
		storeTotal += value
		time.Sleep(time.Second)
	}

	fmt.Println("Total:", ToCurrency(storeTotal))
}

// This file defines methods that operate on the type aliases created in the product.go file.
// Methods can be defined only on types that are created in the same package, which means that I can’t define a method for the []*Product type,
// for example, but I can create an alias for that type and use the alias as the method receiver.

func (group ProductGroup) TotalPrice(category string, resultChannel chan<- float64) {
	var total float64
	for _, p := range group {
		fmt.Println(category, "product:", p.Name)
		total += p.Price
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println(category, "channel sending", ToCurrency(total))
	resultChannel <- total
	fmt.Println(category, "channel send complete")
	// This statement sends the total value through the resultChannel channel, which makes it available to be received elsewhere in the application.
}
