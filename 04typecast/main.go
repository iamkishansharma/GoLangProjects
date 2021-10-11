package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("----- Pizza Rating App -----")
	fmt.Println("Rate from 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate the pizza you had:")
	input, _ := reader.ReadString('\n')

	fmt.Println("Thank you! for rating our Pizza.", input)
	numberRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to rating.")
		numberRating = numberRating + 1
		fmt.Println("Now the rating is: ", numberRating)

	}

}
