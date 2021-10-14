package main

import "fmt"

/*
NOTE:
 1. Nested functon is not supported in GO
 1. Functions can be without name but has to be called exactly after } using ()
*/
func main() {
	fmt.Println("Welcome to FUNCTION")

	// function call
	greet("Morning!")

	a, b := 2, 3
	result := add(a, b)
	fmt.Println("Added using function: ", result)

	result = proAdder(a, b, 2, 3)
	fmt.Println("Added using PRO function: ", result)
}

// func accepting 2 args
// and retuning an int value
func add(aa int, bb int) int {
	return aa + bb
}

// func accepting multiple values as args
// and retuning an int value
func proAdder(props ...int) int {
	total := 0
	for _, val := range props {
		total += val
	}
	return total
}

// function declaration & init with single arg
/// and returns nothing
func greet(param string) {
	fmt.Println("Good", param)
}

// func (){
// 	fmt.Println("Jay shree ram")
// }()
