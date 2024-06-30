package main

import (
	"composition/store"
	"fmt"
)

func UsingTheBoatStruct() {
	boats := []*store.Boat{
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 650.25, 2, true),
	}
	for _, b := range boats {
		fmt.Println("Conventional:", b.Product.Name, "Direct:", b.Name)
	}
}

// func AccessingNestedFieldDirectly() {
// 	rentals := []*store.RentalBoat{
// 		store.NewRentalBoat("Rubber Ring", 10, 1, false, false),
// 		store.NewRentalBoat("Yacht", 50000, 5, true, true),
// 		store.NewRentalBoat("Super Yacht", 100000, 15, true, true),
// 	}
// 	for _, r := range rentals {
// 		fmt.Println("Rental Boat:", r.Name, "Rental Price:", r.Price(0.2))
// 	}
// }

func UsingPromotedFields() {
	rentals := []*store.RentalBoat{
		store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
		store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
		store.NewRentalBoat("Super Yacht", 100000, 15, true, true,
			"Dora", "Charlie"),
	}
	for _, r := range rentals {
		fmt.Println("Rental Boat:", r.Name, "Rental Price:", r.Price(0.2),
			"Captain:", r.Captain)
	}
}

func DemonstateHowPromotionIsHandled() {
	product := store.NewProduct("Kayak", "Watersports", 279)
	deal := store.NewSpecialDeal("Weekend Special", product, 50)
	Name, price, Price := deal.GetDetails()
	fmt.Println("Name:", Name)
	fmt.Println("Price field:", price)
	fmt.Println("Price method:", Price)
}

func UsingAnInterface() {
	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball":  store.NewProduct("Soccer Ball", "Soccer", 19.50),
	}
	for key, p := range products {
		fmt.Println("Key:", key, "Price:", p.Price(0.2))
	}
}

func UsingSeparateCaseStatement() {
	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball":  store.NewProduct("Soccer Ball", "Soccer", 19.50),
	}
	for key, p := range products {
		switch item := p.(type) {
		case *store.Product:
			fmt.Println("Name:", item.Name, "Category:", item.Category,
				"Price:", item.Price(0.2))
		case *store.Boat:
			fmt.Println("Name:", item.Name, "Category:", item.Category,
				"Price:", item.Price(0.2))
		default:
			fmt.Println("Key:", key, "Price:", p.Price(0.2))
		}
	}
}

func UsingDescribableInterface() {
	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball":  store.NewProduct("Soccer Ball", "Soccer", 19.50),
	}
	for key, p := range products {
		switch item := p.(type) {
		case store.Describable:
			fmt.Println("Name:", item.GetName(), "Category:", item.GetCategory(),
				"Price:", item.(store.ItemForSale).Price(0.2))
		default:
			fmt.Println("Key:", key, "Price:", p.Price(0.2))
		}
	}
}

func main() {
	// UsingTheBoatStruct()
	// DemonstateHowPromotionIsHandled()
	// UsingAnInterface()
	UsingSeparateCaseStatement()
}
