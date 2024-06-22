package main

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

func WorkingWithMaths() {
	val1 := 279.00
	val2 := 48.95
	Printfln("Abs: %v", math.Abs(val1))
	Printfln("Ceil: %v", math.Ceil(val2))
	Printfln("Copysign: %v", math.Copysign(val1, -5))
	Printfln("Floor: %v", math.Floor(val2))
	Printfln("Max: %v", math.Max(val1, val2))
	Printfln("Min: %v", math.Min(val1, val2))
	Printfln("Mod: %v", math.Mod(val1, val2))
	Printfln("Pow: %v", math.Pow(val1, 2))
	Printfln("Round: %v", math.Round(val2))
	Printfln("RoundToEven: %v", math.RoundToEven(val2))
}

func WorkingWithRandomNumbers() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 5; i++ {
		Printfln("Value %v : %v", i, rnd.Int())
	}
}

// Generating a Random Number Within a Specific Range

func RandomNumberWithSpecifiedLimit() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 5; i++ {
		Printfln("Value %v: %v", i, rnd.Intn(10))
	}
}

func IntRange(min, max int) int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	return rnd.Intn(max-min) + min
}

func RandomNumberWithGivenRange() {
	for i := 0; i < 5; i++ {
		Printfln("Value %v : %v", i, IntRange(10, 20))
	}
}

// Shuffling elements

// The Shuffle function is used to randomly reorder elements, which it does with the use of a custom function

var names = []string{"Alice", "Bob", "Charlie", "Dora", "Edith"}

func ShufflingOrder() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	rnd.Shuffle(len(names), func(first, second int) {
		names[first], names[second] = names[second], names[first]
	})
	for i, name := range names {
		Printfln("Index %v: Name: %v", i, name)
	}
}

// The arguments to the Shuffle function are the number of elements and a function that will swap two elements, which are identified by index.
// The function is called to swap elements randomly. The anonymous function switches two elements in the names slice,
// which means that the use of the Shuffle function has the effect of shuffling the order of the names values.

// Sorting Data

func SortingData() {
	ints := []int{9, 4, 2, -1, 10}
	Printfln("Ints: %v", ints)
	sort.Ints(ints)
	Printfln("Ints Sorted: %v", ints)
	floats := []float64{279, 48.95, 19.50}
	Printfln("Floats: %v", floats)
	sort.Float64s(floats)
	Printfln("Floats Sorted: %v", floats)
	strings := []string{"Kayak", "Lifejacket", "Stadium"}
	Printfln("Strings: %v", strings)
	if !sort.StringsAreSorted(strings) {
		sort.Strings(strings)
		Printfln("Strings Sorted: %v", strings)
	} else {
		Printfln("Strings Already Sorted: %v", strings)
	}
}

// the functions sort the elements in place, rather than creating a new slice.
// If you want to create a new, sorted slice, then you must use the built-in make and copy functions

func NotMakingInPlaceSort() {
	ints := []int{9, 4, 2, -1, 10}
	sortedInts := make([]int, len(ints))
	copy(sortedInts, ints)
	sort.Ints(sortedInts)
	Printfln("Ints: %v", ints)
	Printfln("Ints Sorted: %v", sortedInts)
}

// Searching Sorted Data

func SearchingSortedData() {
	ints := []int{9, 4, 2, -1, 10}
	sortedInts := make([]int, len(ints))
	copy(sortedInts, ints)
	sort.Ints(sortedInts)
	Printfln("Ints: %v", ints)
	Printfln("Ints Sorted: %v", sortedInts)
	indexOf4 := sort.SearchInts(sortedInts, 4)
	indexOf3 := sort.SearchInts(sortedInts, 3)
	Printfln("Index of 4: %v", indexOf4)
	Printfln("Index of 3: %v", indexOf3)
}

// searching for a value that is in the slice produces the same result as a search for a nonexistent value:

// Result of above function when executed

// Ints: [9 4 2 -1 10]
// Ints Sorted: [-1 2 4 9 10]
// Index of 4: 2
// Index of 3: 2

func VerifySearchedData() {
	ints := []int{9, 4, 2, -1, 10}
	sortedInts := make([]int, len(ints))
	copy(sortedInts, ints)
	sort.Ints(sortedInts)
	Printfln("Ints: %v", ints)
	Printfln("Ints Sorted: %v", sortedInts)
	indexOf4 := sort.SearchInts(sortedInts, 4)
	indexOf3 := sort.SearchInts(sortedInts, 3)
	Printfln("Index of 4: %v (present: %v)", indexOf4, sortedInts[indexOf4] == 4)
	Printfln("Index of 3: %v (present: %v)", indexOf3, sortedInts[indexOf3] == 3)
}

// Sorting Custon Data

// See the implementation in productsort.go file

func SortingCustomData() {
	products := []Product{
		{"Kayak", 279},
		{"Lifejacket", 49.95},
		{"Soccer Ball", 19.50},
	}
	ProductSlices(products)
	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}
}

// Specifying the comparision function

func SpecifyingTheComparisionFunction() {
	products := []Product{
		{"Lifejacket", 49.95},
		{"Kayak", 279},
		{"Soccer Ball", 19.50},
	}
	SortWith(products, func(p1, p2 Product) bool {
		return p1.Name > p2.Name
	})
	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}
}

func main() {
	// Printfln("Hello, Math and Sorting")
	// WorkingWithMaths()
	// WorkingWithRandomNumbers()
	// RandomNumberWithSpecifiedLimit()
	// RandomNumberWithGivenRange()
	// ShufflingOrder()
	// SortingData()
	// NotMakingInPlaceSort()
	// SearchingSortedData()
	// VerifySearchedData()
	// SortingCustomData()
	SpecifyingTheComparisionFunction()
}
