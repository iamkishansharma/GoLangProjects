package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Wecome to User Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number: ")

	/// comma ok
	// OR comma eror syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("You have entered ", input)
	fmt.Printf("Type of this var is: %T \n", input)

}
