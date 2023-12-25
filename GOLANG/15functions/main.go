package main

import "fmt"

func main() {
	greeter("functions")
	result := adder(3, 5)
	fmt.Println("Result is : ", result)

	proRes, myMessage := proAdder(2, 4, 5, 3, 2)

	fmt.Printf("ProResult is : %v , ProMessage is : %v", proRes, myMessage)
}

func greeter(topic string) {
	fmt.Println("Welcome to lecture : ", topic)
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	totalValue := 0

	// Here we are using comma ok syntax because there is no use of index
	for _, val := range values {
		totalValue += val
	}

	return totalValue, "Hello from proAdder"
}
