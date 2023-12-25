package main

import "fmt"

func main() {
	loginCount := 9

	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "something else"
	} else {
		result = "Exactly equal 10"
	}

	fmt.Println(result)

	// There is a special syntax

	if num := 20; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is NOT less than 10")
	}

	// if err != nil{

	// }
}
