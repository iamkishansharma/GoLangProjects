package main

import (
	"fmt"
	"net/url"
)

const URL = "https://lco.dev/tutcss/?fname=Kaali&lname=Sharma&fav_language=HTML&birthday=2021-10-12"

func main() {
	fmt.Println("Welcome to Url Handling")
	fmt.Println("My url: ", URL)

	// parsing
	result, err := url.Parse(URL)
	checkError(err)

	// fmt.Println("Scheme: ", result.Scheme)
	// fmt.Println("Host: ", result.Host)
	// fmt.Println("Scheme: ", result.Scheme)
	// fmt.Println("Path: ", result.Path)
	// fmt.Println("Port: ", result.Port())
	fmt.Println("RawQuery: ", result.RawQuery)

	qparams := result.Query()

	fmt.Printf("Type of QueryParams: %T\n", qparams)
	fmt.Printf("Query result: %s %s\n", qparams["fname"], qparams["lname"])

	// another way of creating URL
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	newUrl := partsOfUrl.String()

	fmt.Println(newUrl)

}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
