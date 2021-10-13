package main

import (
	"fmt"
)

// also known as HashTable in some languages
func main() {
	fmt.Println("welcome to MAPS")

	languages := make(map[string]string)

	languages["en"] = "English"
	languages["hi"] = "Hindi"
	languages["sp"] = "Spanish"
	languages["np"] = "Nepali"

	// map[en:English hi:Hindi np:Nepali sp:Spanish]
	fmt.Println(languages)

	// deleting value
	delete(languages, "hi")
	fmt.Println(languages)

	fmt.Println("-------- Another Way -------")
	fruits := map[int]string{2: "Apple", 5: "Banana", 3: "Orange"}
	fmt.Println(fruits)

	fmt.Println("-------- for Loop in MAP -------")
	for key, value := range languages {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
