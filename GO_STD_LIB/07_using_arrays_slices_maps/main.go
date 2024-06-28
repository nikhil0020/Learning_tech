package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

// Working With Arrays

func WorkingWithArrays() {
	var names [3]string
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println(names)
}

// Using the Array Literal Syntax

func UsingArrayLiteralSyntax() {
	names := [3]string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names)
}

// CREATING MULTIDIMENSIONAL ARRAYS
// Go arrays are one-dimensional but can be combined to create multidimensional arrays, like this:

var coords [3][3]int

// Working with Slices

// The best way to think of slices is as a variable-length array because they
// are useful when you don’t know how many values you need to store or when the number changes over time.
// One way to define a slice is to use the built-in make function

func WorkingWithSlices() {
	names := make([]string, 3)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println(names)
}

func AppendingElementToSlice() {
	names := []string{"Kayak", "Lifejacket", "Paddle"}
	names = append(names, "Hat", "Gloves")
	fmt.Println(names)
}

func AllocatingAdditionalSpaceToSlice() {
	names := make([]string, 3, 6)
	// 3 is length
	// 6 is capacity
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println("len:", len(names))
	fmt.Println("cap:", cap(names))
}

func AddingElementToSlice() {
	names := make([]string, 3, 6)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	appendedNames := append(names, "Hat", "Gloves")
	names[0] = "Canoe"
	fmt.Println("names:", names)
	fmt.Println("appendedNames:", appendedNames)
}

// Appending One Slice to Another

// The second argument is followed by three periods (...), which is required because
// the built-in append function defines a variadic parameter it is enough to know that
// you can append the contents of one slice to another slice, just as long as the three periods are used

func AppendingOneSliceToAnother() {
	names := make([]string, 3, 6)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	moreNames := []string{"Hat Gloves"}
	appendedNames := append(names, moreNames...)
	fmt.Println("appendedNames:", appendedNames)
}

// Creating Slices From Existing Arrays

// Ranges are expressed within square brackets, with the low and high values separated
// by a colon. The first index in the slice is set to be the low value, and the length
// is the result of the high value minus the low value. This means that the range [1:3]
// creates a range whose zero index is mapped into index 1 of the array and whose length is 2.

func CreatingSlicesFromExistingArrays() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3]
	allNames := products[:]
	fmt.Println("someNames:", someNames)
	fmt.Println("allNames", allNames)
}

// Specifying Capacity When Creating a Slice from an Array

func SpecifyingCapacityWhenCreatingSliceFromArray() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3:3]
	// 1 is lower limit
	// 3 is higher limit
	// last 3 is maximum
	allNames := products[:]
	someNames = append(someNames, "Gloves")
	//someNames = append(someNames, "Boots")
	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames", allNames)
	fmt.Println("allNames len", len(allNames), "cap:", cap(allNames))
}

// Creating Slices from Other Slices

func CreatingSlicesFromOtherSlices() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames := products[1:]
	someNames := allNames[1:3]
	allNames = append(allNames, "Gloves")
	allNames[1] = "Canoe"
	fmt.Println("someNames:", someNames)
	fmt.Println("allNames", allNames)
}

// Using the copy Function

// The copy function is used to copy elements between slices. This function can be used to ensure
// that slices have separate arrays and to create slices that combine elements from different sources.

// Using the copy Function to Ensure Slice Array Separation

// If desination slice length is zero or uninitialized then no element will be copied
// Also additional Capacity is also not considered
func CopyingFromASlice() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}

	allNames := products[:]

	someNames := make([]string, 3, 6)

	copy(someNames, allNames)
	fmt.Println("someNames:", someNames)
	fmt.Println("allNames", allNames)
}

// Specifying Ranges When Copying Slices

// Fine-grained control over the elements that are copied can be achieved using ranges

func CopyingRange() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames := products[:]
	someNames := []string{"Boots", "Canoe"}
	copy(someNames[1:], allNames[2:3])
	// fmt.Println(allNames[0:3])
	fmt.Println("someNames:", someNames)
	fmt.Println("allNames", allNames)
}

// Deleting Slice Elements

func DeletingSliceElements() {
	products := [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	deleted := append(products[:2], products[3:]...)
	fmt.Println("Deleted:", deleted)
}

// Enumerating Slices
// Slices are enumerated in the same way as arrays, with the for and range keywords

func EnumeratingSlices() {
	products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	for index, value := range products[2:] {
		fmt.Println("Index:", index, "Value:", value)
	}
}

// Sorting Slices
// There is no built-in support for sorting slices, but the standard library includes the sort package,

func SortingSlice() {
	products := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	sort.Strings(products)
	for index, value := range products {
		fmt.Println("Index:", index, "Value:", value)
	}
}

// Comparing Slices
// Go restricts the use of the comparison operator so that slices can be compared only to the nil value.
// Comparing two slices produces an error,

// There is one way to compare slices, however. The standard library includes a package named reflect,
// which includes a convenience function named DeepEqual

func ComparingSlicesWithReflect() {
	p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	p2 := p1
	fmt.Println("Equal:", reflect.DeepEqual(p1, p2))
}

// Getting the array underlying the slices

func GettingTheArrayUnderlyingSlice() {
	p1 := []string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	arrayPtr := (*[3]string)(p1)
	array := *arrayPtr
	fmt.Println(array)
}

func WorkingWithMaps() {
	productsMap := make(map[string]float64, 10)
	productsMap["Kayak"] = 279
	productsMap["Lifejacket"] = 48.95
	fmt.Println("Map size:", len(productsMap))
	fmt.Println("Price:", productsMap["Kayak"])
	fmt.Println("Price:", productsMap["Hat"])
}

// Using the Map Literal Syntax

func UsingTheMapLiteralSyntax() {
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
	}
	fmt.Println("Map size:", len(products))
	fmt.Println("Price:", products["Kayak"])
	fmt.Println("Price:", products["Hat"])
}

// Checking for Items in a Map

// The problem with this code is that products["Hat"] returns zero, but it isn’t known whether this
// is because zero is the stored value or because there is no value associated with the key Hat.
// To solve this problem, maps produce two values when reading a value
func CheckingForItemsInMap() {
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	// value, ok := products["Hat"]
	// if ok {
	// 	fmt.Println("Stored value:", value)
	// } else {
	// 	fmt.Println("No stored value")
	// }

	// Another ways , often used
	if value, ok := products["Hat"]; ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}
}

// Removing Items from a Map

func RemovingItemsFromMap() {
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	delete(products, "Hat")
	if value, ok := products["Hat"]; ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}
}

// Enumerating the Contents of a Map

// Maps are enumerated using the for and range keywords

func EnumeratingContentsOfMap() {
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	for key, value := range products {
		fmt.Println("Key:", key, "Value:", value)
	}
}

// Enumerating a Map in Order

// If you want to get the values in a map in order,
// then the best approach is to
// enumerate the map and
// create a slice containing the keys,
// sort the slice, and then enumerate the slice to read the values from the map

func EnumerateMapInOrder() {
	products := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}

	keys := make([]string, 0, len(products))

	for key, _ := range products {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println("Key: ", key, "Value: ", products[key])
	}
}

// Understanding the dual nature of strings

// Indexing and slicing string

func IndexingSlicingString() {
	var price string = "$48.95"
	var currency byte = price[0]
	var amountString string = price[1:]
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	fmt.Println("Currency:", currency)
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}
}

// Converting the result

func ConvertingTheResult() {
	var price string = "$48.95"
	var currency string = string(price[0])
	var amountString string = price[1:]
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	fmt.Println("Currency:", currency)
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}
}

// Changing the currency symbol

func ChangingTheCurrencySymbol() {
	var price string = "€48.95"
	var currency string = string(price[0])
	var amountString string = price[1:]
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	fmt.Println("Currency:", currency)
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}
}

// Obtaining the length of string
// here we are demonstrating that euro is taking 3 bytes not one

func ObtainingTheLengthOfString() {
	var price string = "€48.95"
	var currency string = string(price[0])
	var amountString string = price[1:]
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	fmt.Println("Length:", len(price))
	fmt.Println("Currency:", currency)
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}
}

// Output of above code
// Length: 8
// Currency: â
// Parse Error: strconv.ParseFloat: parsing "\x82\xac48.95": invalid syntax

// Here we can see that length is 8 instread of 6

// Converting to rune

func ConvertingToRune() {
	var price []rune = []rune("€48.95")
	var currency string = string(price[0])
	var amountString string = string(price[1:])
	amount, parseErr := strconv.ParseFloat(amountString, 64)
	fmt.Println("Length:", len(price))
	fmt.Println("Currency:", currency)
	if parseErr == nil {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse Error:", parseErr)
	}
	// the rune type is an alias for int32, which means that printing out a rune value will
	// display the numeric value used to represent the character. This means, as with the byte example previously,
	// I have to perform an explicit conversion of a single rune into a string
}

// Enumerating Strings

// A for loop can be used to enumerate the contents of a string.

func EnumeratingString() {
	var price = "€48.95"
	for index, char := range price {
		fmt.Println(index, char, string(char))
	}
}

// Enumerating Bytes in a String

func EnumeratingBytesInString() {
	var price = "€48.95"
	for index, char := range []byte(price) {
		fmt.Println(index, char)
	}
}

// The index values are sequential, and the values of individual bytes are displayed
// without being interpreted as parts of the characters they represent.

func main() {
	// fmt.Println("Hello, Collections")
	coords[0][2] = 10
	// WorkingWithArrays()
	// UsingArrayLiteralSyntax()
	// WorkingWithSlices()
	// AppendingElementToSlice()
	// AllocatingAdditionalSpaceToSlice()
	// AddingElementToSlice()
	// AppendingOneSliceToAnother()
	// CreatingSlicesFromExistingArrays()
	// SpecifyingCapacityWhenCreatingSliceFromArray()
	// CreatingSlicesFromOtherSlices()
	// CopyingFromASlice()
	// CopyingRange()
	// DeletingSliceElements()
	// EnumeratingSlices()
	// SortingSlice()
	// ComparingSlicesWithReflect()
	// GettingTheArrayUnderlyingSlice()
	// WorkingWithMaps()
	// UsingTheMapLiteralSyntax()
	// CheckingForItemsInMap()
	// RemovingItemsFromMap()
	// EnumeratingContentsOfMap()
	// EnumerateMapInOrder()
	// IndexingSlicingString()
	// ConvertingTheResult()
	// ChangingTheCurrencySymbol()
	// ObtainingTheLengthOfString()
	// ConvertingToRune()
	// EnumeratingString()
	EnumeratingBytesInString()
}
