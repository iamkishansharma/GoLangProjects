package main

import "fmt"

func main() {
	fmt.Println("Welcome to METHODs")

	kishan := User{
		123,
		"Kishan Sharma",
		"kishan@email.com",
		true,
	}
	fmt.Println("User: ", kishan.name)
	fmt.Println("Getting user verification status....")
	kishan.GetStatus()
	kishan.NewMail()
	fmt.Println("Original Email: ", kishan.email)

}

type User struct {
	id        int
	name      string
	email     string
	isVerifed bool
}

func (u User) GetStatus() {
	fmt.Println("Is user verified: ", u.isVerifed)
}

func (u User) NewMail() {
	u.email = "abcd@gmail.com"
	fmt.Println("Updated Email: ", u.email)
}
