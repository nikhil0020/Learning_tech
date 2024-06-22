package main

import "fmt"

// Formatting strings

// The Println and Fprintln functions add spaces between all the values,
// but the Print and Fprint functions only add spaces between values that are not strings

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

func getProductName(index int) (name string, err error) {
	if len(Products) > index {
		name = fmt.Sprintf("Name of the product: %v", Products[index].Name)
	} else {
		err = fmt.Errorf("error for index %v", index)
	}
	return
}

func UsingFormattedStrings() {
	name, _ := getProductName(1)
	fmt.Println(name)

	_, err := getProductName(10)
	fmt.Println(err.Error())
}

// Using the general purpose formatting verbs

func UsingGenPurposeFormattingVerbs() {
	Printfln("Value: %v", Kayak)
	Printfln("Go syntax: %#v", Kayak)
	Printfln("Type: %T", Kayak)
}

// Controlling struct formatting
// Go has a default format for all data types that the %v verb relies on.
// For structs, the default value lists the field values within curly braces.
// The default verb can be modified with a plus sign to include the field names

func ControllingStructFormatting() {
	Printfln("Value with fields: %+v", Kayak)
}

// The fmt package supports custom struct formatting through an interface named Stringer that is defined as follows:

// type Stringer interface {
// 	String() string
// }

// The String method specified by the Stringer interface will be used to obtain a string representation of any type that defines it

// Check product.go to see String method implementation

// The String method will be invoked automatically when a string representation of a Product value is required.

func UsingTheIntegerFormattingVerbs() {
	number := 250
	Printfln("Binary: %b", number)
	Printfln("Decimal: %d", number)
	Printfln("Octal: %o, %O", number, number)
	Printfln("Hexadecimal: %x, %X", number, number)
}

func UsingFloatingPointFormattingVerbs() {
	number := 279.00
	Printfln("Decimalless with exponent: %b", number)
	Printfln("Decimal with exponent: %e", number)
	Printfln("Decimal without exponent: %.2f", number)
	Printfln("Hexadecimal: %x, %X", number, number)
}

func UsingFormattingVerbModifier() {
	number := 279.00
	Printfln("Sign: >>%+.2f<<", number)
	Printfln("Zeros for Padding: >>%010.2f<<", number)
	Printfln("Right Padding: >>%-8.2f<<", number)
}

func UsingStringAndCharacterFormattingVerbs() {
	name := "Kayak"
	Printfln("String: %s", name)
	Printfln("Character: %c", []rune(name)[0])
	Printfln("Unicode: %U", []rune(name)[0])
}

func UsingBooleanFormattingVerbs() {
	name := "Kayak"
	Printfln("Bool: %t", len(name) > 1)
	Printfln("Bool: %t", len(name) > 100)
}

func UsingPointerFormattingVerbs() {
	name := "Kayak"
	Printfln("Pointer: %p", &name)
}

func ScanningString() {
	var name string
	var category string
	var price float64
	fmt.Print("Enter text to scan: ")
	n, err := fmt.Scan(&name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func DealingWithNewlineCharacters() {
	var name string
	var category string
	var price float64
	fmt.Print("Enter text to scan: ")
	n, err := fmt.Scanln(&name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func UsingDifferentStringSource() {
	var name string
	var category string
	var price float64
	source := "Lifejacket Watersports 48.95"
	n, err := fmt.Sscan(source, &name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func UsingScanningTemplate() {
	var name string
	var category string
	var price float64
	source := "Product Lifejacket Watersports 48.95"
	template := "Product %s %s %f"
	n, err := fmt.Sscanf(source, template, &name, &category, &price)
	if err == nil {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func main() {
	// fmt.Println("Product:", Kayak.Name, "Price:", Kayak.Price)
	// UsingFormattedStrings()
	// UsingGenPurposeFormattingVerbs()
	// ControllingStructFormatting()
	// UsingTheIntegerFormattingVerbs()
	// UsingFloatingPointFormattingVerbs()
	// UsingFormattingVerbModifier()
	// UsingStringAndCharacterFormattingVerbs()
	// UsingBooleanFormattingVerbs()
	// UsingPointerFormattingVerbs()
	// ScanningString()
	// DealingWithNewlineCharacters()
	// UsingDifferentStringSource()
	UsingScanningTemplate()
}
