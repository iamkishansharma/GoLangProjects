package main

import "fmt"

func main() {
	fmt.Println("Welcome to POINTERS")

	// declairation of pointer variable
	var ptr *int
	fmt.Println("Value of pointerOne is: ", ptr)

	num := 25
	// declairation & init of pointer variable
	pointerTwo := &num
	fmt.Println("Address of num is: ", &num)
	fmt.Println("pointerTwo is pointing at: ", pointerTwo)
	fmt.Println("Extracting value from that address: ", *pointerTwo)

	// Adding 5 to that value where aout pointer is pointing at
	*pointerTwo += 5
	fmt.Println("Add 5 to pointerTwo Pointing at: ", *pointerTwo)
	fmt.Println("Now pointerTwo is pointing at: : ", pointerTwo)
	fmt.Println("Now Address of num is: ", &num)
	fmt.Println("Now Value of num is: ", num)

}
