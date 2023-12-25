package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

type FoodNameGetter func(string) string

type Food struct {
	name   string
	getter FoodNameGetter // declare function
}

type Service struct {
	name string
	user User
}

func main() {
	// * No inheritance in golang , No super or parent

	// ! Wrong way to initialize structs is A := User(values with comma)
	// * Right way to initialize structs is A := User{values with comma}

	nikhil := User{"Nikhil", "gmail.com", true, 21}

	fmt.Println(nikhil)

	// * We can use . operator to access structs attributes and members

	fmt.Println(nikhil.Age)
	// * For detailed version of structs we can do this

	fmt.Printf("Nikhil details are : %+v\n", nikhil)

	// Functions as a field

	pizza := Food{
		name: "Pizza",
		getter: func(name string) string { // declare function body
			return "This is " + name + "."
		},
	}

	fmt.Println(pizza.getter(pizza.name))

	// Now, for a nested struct, we can do like this:

	google := Service{
		name: "Google",
		user: User{
			Name:   "Nikhil",
			Email:  "abcd@gmail.com",
			Status: false,
			Age:    21,
		},
	}

	facebook := Service{
		name: "Facebook",
		user: User{
			Name:   "Ram",
			Email:  "Ram@siya.in",
			Status: false,
			Age:    100000,
		},
	}

	fmt.Printf("Services used is %+v\n", google)

	// Slices of structs

	var services = []Service{google}

	services = append(services, facebook)

	fmt.Println(services)
}
