package main

import "fmt"

func main() {
	defer fmt.Printf(" Hello ")
	fmt.Printf(" World ")
	defer fmt.Printf(" From ")
	defer fmt.Printf(" Nikhil")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
