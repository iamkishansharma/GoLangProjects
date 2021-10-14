package main

import "fmt"

/*
1. LIFO - last In First Out execution
2. We can say it executes in reverse order
*/
func main() {
	fmt.Println("Welcome to DEFER")
	defer fmt.Println("3")
	defer fmt.Println("2")
	defer fmt.Println("1")

	defer fmt.Println("Hello, World!")
	callDefer()
}

func callDefer() {
	for i := 0; i < 11; i++ {
		defer fmt.Println("Defer: ", i)
	}
}

/*
output:
Welcome to DEFER
Defer:  10
Defer:  9
Defer:  8
Defer:  7
Defer:  6
Defer:  5
Defer:  4
Defer:  3
Defer:  2
Defer:  1
Defer:  0
Hello, World!
1
2
3
*/
