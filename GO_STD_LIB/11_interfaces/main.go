package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

// we are using pointer type for Product
func UsingAnInterface() {
	expenses := []Expense{
		&Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 89.50, []string{}},
	}

	for _, expense := range expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}
}

// Using Interface in a function

func calcTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total += item.getCost(true)
	}
	return
}

// Using an interface for struct fields

type Account struct {
	AccountNo int
	expenses  []Expense
}

func UsingAnInterfaceForStructFields() {
	account := Account{
		AccountNo: 12345,
		expenses: []Expense{
			&Product{"Kayak", "Watersports", 275},
			Service{"Boat Cover", 12, 89.5, []string{}},
		},
	}

	for _, expense := range account.expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}

	fmt.Println("Total:", calcTotal(account.expenses))
}

// Understanding the effect of pointer method receivers

// func ValueReceiver() {
// 	product := Product{"Kayak", "Watersports", 275}
// 	var expense Expense = product
// 	product.price = 100
// 	fmt.Println("Product field value:", product.price)
// 	fmt.Println("Expense method result:", expense.getCost(false))
// }

// Output of value

// Product field value: 100
// Expense method result: 275

// The Product value was copied when it was assigned to the Expense variable,
// which means that the change to the price field does not affect the result from the getCost method.

func PointerReceiver() {
	product := Product{"Kayak", "Watersports", 200}
	var expense Expense = &product

	product.price = 500

	fmt.Println("Product field value:", product.price)
	fmt.Println("Expense method result:", expense.getCost(false))
}

// Comparing interface values

func ComparingInterfaceValues() {
	var e1 Expense = &Product{name: "Kayak"}
	var e2 Expense = &Product{name: "Kayak"}
	var e3 Expense = Service{description: "Boat Cover"}
	var e4 Expense = Service{description: "Boat Cover"}
	fmt.Println("e1 == e2", e1 == e2)
	fmt.Println("e3 == e4", e3 == e4)
}

// The first two Expense values are not equal. That’s because the dynamic type for these
// values is a pointer type, and pointers are equal only if they point to the same memory location. The second two
// Expense values are equal because they are simple struct values with the same field values.

// Interface equality checks can also cause runtime errors if the dynamic type is not comparable.
// since i added features in Service struct, it will give error while comparing e3 with e4

//
//
// Performing Type Assertion

// Interfaces can be useful, but they can present problems, and it is often useful to be able to
// access the dynamic type directly, which is known as type narrowing,
// the process of moving from a less precise type to a more precise type.

func TypeAssertion() {
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.50, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		// &Product{"Kayak", "Watersports", 275},
		// above line will give error
		// panic: interface conversion: main.Expense is *main.Product, not main.Service
	}

	for _, expense := range expenses {
		s := expense.(Service)
		fmt.Println("Service:", s.description, "Price:", s.monthlyFee*float64(s.durationMonths))
	}
}

// The Go runtime has tried to perform the assertion and failed. To avoid this issue,
// there is a special form of type assertion that indicates whether an assertion could be performed

func TypeAssertionWithSpecialForm() {
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.50, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		&Product{"Kayak", "Watersports", 275},
	}

	for _, expense := range expenses {
		if s, ok := expense.(Service); ok {
			fmt.Println("Service:", s.description, "Price:", s.monthlyFee*float64(s.durationMonths))
		} else {
			fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
		}
	}
}

// Switching on Dynamic types

func SwitchingOnDynamicTypes() {
	expenses := []Expense{
		Service{"Boat Cover", 12, 89.50, []string{}},
		Service{"Paddle Protect", 12, 8, []string{}},
		&Product{"Kayak", "Watersports", 275},
	}

	for _, expense := range expenses {
		switch value := expense.(type) {
		case Service:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
		case *Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		default:
			fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))

		}
	}
}

// Using Empty Interface

type Person struct {
	name, city string
}

func UsingEmptyInterface() {
	var expense Expense = &Product{"Kayak", "Watersports", 275}
	data := []interface{}{
		expense,
		Product{"Lifejacket", "Watersports", 48.95},
		Service{"Boat Cover", 12, 89.50, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}

	for _, item := range data {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		case *Product:
			fmt.Println("Product Pointer:", value.name, "Price:", value.price)
		case Service:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
		case Person:
			fmt.Println("Person:", value.name, "City:", value.city)
		case *Person:
			fmt.Println("Person pointer", value.name, "City:", value.city)
		case string, bool, int:
			fmt.Println("Built-in types:", value)
		default:
			fmt.Println("Default:", value)
		}
	}
}

// The empty interface can also be used for variadic parameters, which allows a
// function to be called with any number of arguments, each of which can be any type

func processItems(items ...interface{}) {
	for _, item := range items {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		case *Product:
			fmt.Println("Product Pointer:", value.name, "Price:", value.price)
		case Service:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
		case Person:
			fmt.Println("Person:", value.name, "City:", value.city)
		case *Person:
			fmt.Println("Person pointer", value.name, "City:", value.city)
		case string, bool, int:
			fmt.Println("Built-in types:", value)
		default:
			fmt.Println("Default:", value)
		}
	}
}

func UsingEmptyInterfaceWithVariadicParameter() {
	var expense Expense = &Product{"Kayak", "Watersports", 275}
	var data []interface{} = []interface{}{
		expense,
		Product{"Lifejacket", "Watersports", 48.95},
		Service{"Boat Cover", 12, 89.50, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}

	processItems(data...)
}

func main() {
	// kayak := Product{"Kayak", "WaterSport", 275}
	// insurance := Service{"Boat Cover", 12, 89.5}
	// fmt.Println("Product:", kayak.name, "Price:", kayak.price)
	// fmt.Println("Service:", insurance.description, "Price:", insurance.monthlyFee*float64(insurance.durationMonths))
	// UsingAnInterface()
	// UsingAnInterfaceForStructFields()
	// ValueReceiver()
	// PointerReceiver()
	// ComparingInterfaceValues()
	// TypeAssertion()
	// TypeAssertionWithSpecialForm()
	// SwitchingOnDynamicTypes(s)
	// UsingEmptyInterface()
	UsingEmptyInterfaceWithVariadicParameter()
}
