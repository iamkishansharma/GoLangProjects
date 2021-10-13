package main

import "fmt"

func main() {
	fmt.Println("Welcome to STRUCT")

	kishan := User{
		"Kishan Kr Sharma",
		"kishan@lco.dev",
		true,
		121,
	}
	// only values without key
	fmt.Println("Details of KISHAN:", kishan)
	fmt.Printf("Type of KISHAN is: %T\n", kishan)
	// Printing details with all details
	fmt.Printf("KISHAN details: %+v\n", kishan)

	// printing indidually
	fmt.Printf("Name is: %v\n", kishan.name)
	fmt.Printf("Email is: %v\n", kishan.email)

}

// declairation of structure
type User struct {
	name   string
	email  string
	status bool
	id     int
}
