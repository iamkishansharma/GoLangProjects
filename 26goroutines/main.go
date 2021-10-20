package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // usually they are pointers
var mut sync.Mutex    // locking-unlocking the memory space
func main() {
	fmt.Println("Welcome to Go Routines!")

	websitelist := []string{
		"https://heycode.tech",
		"https://google.com",
		"https://github.com",
		"https://lco.dev",
		"https://fb.com",
		"https://netflix.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}

	// will not allow main function to exit until the work done
	wg.Wait()
	fmt.Println(signals)

	/* go greeter("Hello")
	greeter("World") */
}

func getStatusCode(endpoint string) {
	// i'm done
	defer wg.Done()

	response, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Opps....")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", response.StatusCode, endpoint)
	}

}

/*
func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println(s)
	}
}
*/
