package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to If-Else")

	// One example of If-Else
	fmt.Println("--------- GRADE CHECKER ---------")

	// creating reader for takig input from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your total marks: ")

	// comma ok syntax
	// read string ultil user hits ENTER (or next line)
	input, _ := reader.ReadString('\n')

	// converting string input to float64 for accuracy
	totalMarks, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if totalMarks <= 100 {
		if totalMarks >= 80 && totalMarks <= 100 {
			fmt.Println("Outstanding!, You got: O")
		} else if totalMarks >= 70 && totalMarks < 80 {
			fmt.Println("Congratulations!, You got: A+")
		} else if totalMarks >= 60 && totalMarks < 70 {
			fmt.Println("Yay!, You got: A")
		} else if totalMarks >= 50 && totalMarks < 60 {
			fmt.Println("Good!, You got: B")
		} else if totalMarks >= 40 && totalMarks < 50 {
			fmt.Println("Good!, You got: c")
		} else {
			fmt.Println("Sorry!, You are failed. :(")
		}
	} else {
		fmt.Println("Wrong input!, Marks cannot cross 100.")
		return
	}

	fmt.Println("----- Odd Even Checker -----")
	// Two Example of If-Else
	if num := 10; num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

}
