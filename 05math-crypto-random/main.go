package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to Maths in GoLang")
	var numberOne int = 8
	var numberTwo int = 4

	fmt.Println("Addition: ", numberOne+numberTwo)
	fmt.Println("Subtraction: ", numberOne-numberTwo)
	fmt.Println("Multiply: ", numberOne*numberTwo)
	fmt.Println("Division: ", numberOne/numberTwo)

	// Random Number with MATH
	// import math/rand & import time
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println("Random Number: ", rand.Intn(6))

	// Random Number with CRYPTO
	// import crypto/rand
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(11))
	fmt.Println("Random Number: ", randomNumber)
}
