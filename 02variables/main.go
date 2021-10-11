package main

import "fmt"

//NOTE:: Variable name that starts with a capital letter are PUBLIC
const Website = "www.heycode.tech"

func main() {
	var fullname string = "Kishan Kr Sharma"
	fmt.Println(fullname)
	fmt.Printf("This variable is a type of: %T \n--------\n", fullname)

	var age int = 23
	fmt.Println(age)
	fmt.Printf("This variable is a type of: %T \n--------\n", age)

	var isGraduate bool = true
	fmt.Println(isGraduate)
	fmt.Printf("This variable is a type of: %T \n--------\n", isGraduate)

	var smallInt uint8 = 235
	// It will take values from 0 to 255
	fmt.Println(smallInt)
	fmt.Printf("This variable is a type of: %T \n--------\n", smallInt)

	var smallFloat float32 = 235.123456789
	// reades only 5 digits after Decimal
	fmt.Println(smallFloat)
	fmt.Printf("This variable is a type of: %T \n--------\n", smallFloat)

	////////////////// Default values and aliases
	/*
		int & float has 0
		string has nothing
	*/
	var anyVar int
	fmt.Println(anyVar)
	fmt.Printf("This variable is a type of: %T \n--------\n", anyVar)

	////////////// Implicit type
	var movie = "The Lion King"
	fmt.Println(movie)
	fmt.Printf("This variable is a type of: %T \n--------\n", movie)

	////////////// Other way of declairing and defining variable without using VAR
	// NOTE:: The walrus operator is allowed only inside functions not outside of func
	song := "Bohemian Ranspodey"
	fmt.Println(song)
	fmt.Printf("This variable is a type of: %T \n--------\n", song)
}

/*
 Check out more about numeric types here
 https://golang.org/ref/spec#Numeric_types
*/
