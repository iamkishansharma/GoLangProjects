package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to thr TIME")

	presentTime := time.Now()
	fmt.Println("It is : ", presentTime)

	// The format syntax has to be like
	// "01-02-2006 15:04:05 Monday"
	fmt.Println("FormattedTime: ", presentTime.Format("01/02/2006 15:04:05 Monday"))

	//
	createdDate := time.Date(2021, time.August, 02, 10, 20, 03, 0, time.UTC)
	fmt.Println("Created date: ", createdDate.Format("01/02/2006 15:04:05 Monday"))
}
