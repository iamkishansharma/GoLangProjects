package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to CHANNELS in Go")

	myCh := make(chan int, 2) // this 2 allows send 2 elements but receive as user receives

	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5
	wg.Add(2)
	// receive only (-- <- --)
	go func(ch <-chan int, w *sync.WaitGroup) {
		val, isChannelOpen := <-myCh
		fmt.Println("Channel is Open: ", isChannelOpen)
		// if channel open read value
		if isChannelOpen {
			fmt.Println("My Channel have: ", val)
		}
		wg.Done()
	}(myCh, wg)

	// send only
	go func(ch chan<- int, w *sync.WaitGroup) {
		// myCh <- 0
		// close(myCh)
		myCh <- 5
		// myCh <- 6
		wg.Done()

	}(myCh, wg)

	wg.Wait()
}
