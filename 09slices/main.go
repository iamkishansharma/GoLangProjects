package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to SLICES")

	fruits := []string{"Apple", "Banana", "Peach"}
	fmt.Printf("The type of var is: %T\n", fruits)
	fmt.Printf("Fruit list: %s\n", fruits)
	fmt.Println("Length of fruit list: ", len(fruits))

	fmt.Println("-------- Append Value ---------")
	// appending more elements
	fruits = append(fruits, "Litchi", "Mango")
	// Now the value of fruit is
	fmt.Printf("Now the fruit list: %s\n", fruits)
	// deleting elements

	// Slicing these slices
	newVal := fruits[1:3] // assign new list of ["Banana", "Peach"]
	fmt.Println("Slice of that list: ", newVal)

	// slices with Memory management
	highScore := make([]int, 3)
	highScore[0] = 121
	highScore[1] = 825
	highScore[2] = 947
	// highScore[3] = 111 // this will give outof bound error

	// but using append we can add elements  IMPORTANT
	highScore = append(highScore, 555, 888)
	fmt.Println(highScore)

	fmt.Println("-------- Sort Values ---------")
	// Sorts the slices from low to high value
	fmt.Println("Is the Slice sorted?: ", sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println("Now, Slice sorted?: ", sort.IntsAreSorted(highScore))

	// Remove values from SLICES using INDEX n append
	fmt.Println("-------- Remove Value ---------")
	var languages = []string{"C++", "JavaScript", "Python", "Dart", "Perl"}
	fmt.Println(languages)

	index := 2 // remove 3rd language ie "Python"
	languages = append(languages[:index], languages[index+1:]...)
	fmt.Println("Boom!, Python removed.\n", languages)

}
