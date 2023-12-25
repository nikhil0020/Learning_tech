package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	// bufio is the package used for reading buffer

	reader := bufio.NewReader(os.Stdin)

	// * Comma ok || comma error syntax

	input, _ := reader.ReadString('\n')

	// * There is no try catch in golang , So we have to do error handling like this

	// * if we have to deal with error only then we have can use
	// * comma ok like this _ , err := something

	fmt.Printf("input is %T", input)

	// * type of input will always be string

	// So next we will learn about type casting

}
