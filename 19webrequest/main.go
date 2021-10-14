package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://kishansharma.me"

func main() {
	fmt.Println("Welcome to WebRequest")
	res, err := http.Get(url)
	checkError(err)

	fmt.Printf("Type of response: %T\n", res)
	fmt.Println("Status Code: ", res.StatusCode)

	defer res.Body.Close()
	fmt.Println(res)

	dataByte, err := ioutil.ReadAll(res.Body)
	checkError(err)

	content := string(dataByte)
	fmt.Println(content)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
