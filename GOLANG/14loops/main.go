package main

import "fmt"

func main() {
	// ! There is no while loop in golang

	vowels := []string{"A", "E", "I", "O", "U"}

	// one way to use loops
	for d := 0; d < len(vowels); d++ {
		fmt.Printf("%v ", vowels[d])
	}
	fmt.Printf("\n")

	// 2nd way to use loops

	for i := range vowels {
		fmt.Println(vowels[i])
	}

	// 3rd way to use loops

	for index, vowel := range vowels {
		fmt.Printf("index is %v and value is %v\n", index, vowel)
	}

	// For using while loop , we can do something like this

	roogueValue := 1
	for roogueValue < 10 {
		fmt.Println("Value is : ", roogueValue)
		roogueValue++
	}

	// We  can also use goto

	roogueValue = 1
	for roogueValue < 10 {
		fmt.Println("Value is : ", roogueValue)
		if roogueValue == 5 {
			goto name
		}
		roogueValue++
	}

name:
	fmt.Println("Goto is used when roogueValue is 5")
}
