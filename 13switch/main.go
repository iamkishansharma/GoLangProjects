package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to SWITCH")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open.")
	case 2:
		fmt.Println("Move 2 steps ahead.")
	case 3:
		fmt.Println("Move 3 steps ahead.")
	case 4:
		fmt.Println("Move 4 steps ahead.")
	case 5:
		fmt.Println("Move 5 steps ahead.")
	case 6:
		fmt.Println("Move 6 steps ahead and Roll the dice again.")

	default:
		fmt.Println("What is this??????")
	}
}
