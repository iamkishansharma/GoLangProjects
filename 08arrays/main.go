package main

import "fmt"

func main() {
	fmt.Println("Welcome to ARRAY")

	// Array declairation
	var fruits [6]string

	// Array init
	fruits[0], fruits[1], fruits[2], fruits[3], fruits[4] = "Apple", "Banana", "Peach", "Papaya", "Mango"

	fmt.Println("All fruits are: ", fruits)
	// As we have already given the size of array as 6
	// So it will give 6 even is we add 3 elements
	fmt.Println("Length of fruits: ", len(fruits))

	// Another way of doing the same
	lang := [5]string{"C++", "Python", "Java", "Kotlin"}

	fmt.Println("All languages are: ", lang)

	// As we have already given the size of array as 5
	fmt.Println("Length of languages: ", len(lang))
}
