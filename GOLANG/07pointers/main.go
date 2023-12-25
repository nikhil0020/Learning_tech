package main

import "fmt"

func main() {
	var ptr *int

	fmt.Println(ptr)
	// Output : <nil> because it's not initialized

	// var ptr1 *string

	myNumber := 23

	ptr = &myNumber

	*ptr = *ptr + 2

	fmt.Println("New Value is : ", myNumber)
}
