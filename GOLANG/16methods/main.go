package main

import "fmt"

func main() {
	nikhil := User{"Nikhil", "gmail.com", true, 21}

	fmt.Println(nikhil)

	nikhil.GetStatus()
	nikhil.NewEmail()
	fmt.Println(nikhil)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active : ", u.Status)
}

// If we pass Simply then email will now change
// So we have to use pointers for that
func (u *User) NewEmail() {
	u.Email = "test@email.com"
	fmt.Println("New email is ", u.Email)
}
