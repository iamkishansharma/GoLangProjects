package main

import "fmt"

func main() {
	fmt.Println("Welcome to LOOPS")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	fmt.Println("------ ForLoop Style 1 -------")
	// for loop style:1
	for d := 0; d < len(days); d++ {
		fmt.Printf("Index: %d Day: %s\n", d, days[d])
	}

	fmt.Println("------ ForLoop Style 2 -------")
	// for loop style:2
	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("------ ForLoop Style 3 -------")
	for index, day := range days {
		fmt.Printf("Index: %d Day: %s\n", index, day)

	}

	fmt.Println("------ ForLoop Style 4 -------")
	i := 0
	tableOf := 2
	for i < 10 {
		if i == 8 {
			break
		}
		if i == 7 {
			goto hello
		}
		if i == 5 {
			i++
			continue
		}
		fmt.Printf("%d X %d = %d\n", tableOf, i+1, tableOf*(i+1))
		i++
	}

	// lables
hello:
	fmt.Println("Welcome to Lable")
}
