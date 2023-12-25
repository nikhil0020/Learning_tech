package main

import "fmt"

func main() {
	var username string = "nikhil"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T\n", username)

	var isLoggedIn bool = true

	var smallVal uint8 = 255

	fmt.Println(isLoggedIn, smallVal)

	var smallFloat float32 = 255.325235325235

	fmt.Println(smallFloat)

	// default value and some aliases

	var anotherVariable string

	fmt.Println(len(anotherVariable))

	//implicit type

	var website = "google.com"

	// we cannot use Println for printing types
	fmt.Printf("type is %T\n", website)

	// * walrus operator

	name := "nikhil"
	fmt.Println(name)
	// no var used

	// ! we cannot use walrus operator for global variables
	// ! this can only be used inside a function

	const LoginToken string = "asfjkdashfkjasdf"

	// * In the above string we have used first letter capital because its a public variable

}
