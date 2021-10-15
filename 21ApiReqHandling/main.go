package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to API Request Handling")
	URL := "https://kishanapi.herokuapp.com/api/v1/students"

	getRequest(URL)
	postRequest(URL)
}

/* ---------- GET Request ----------- */
func getRequest(myUrl string) {
	response, err := http.Get(myUrl)
	checkError(err)
	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Cont Length: ", response.ContentLength)

	var resPonseString strings.Builder

	data, _ := ioutil.ReadAll(response.Body)
	// fmt.Println("All data: ", string(data))
	byteCount, _ := resPonseString.Write(data)
	fmt.Println(byteCount)
	fmt.Println(resPonseString.String())
}

/* ---------- POST Request ----------- */
func postRequest(myUrl string) {
	// json payload format
	requestBody := strings.NewReader(`
	{
		"name": "Rocket Singh",
    	"email": "rocket.singh@cmail.com",
   	 	"image": "image-link",
    	"dob": "1991-11-12"
	}
	`)

	myUrl2 := myUrl + "/add/kishanra"

	response, err := http.Post(myUrl2, "application/json", requestBody)
	checkError(err)
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	fmt.Println("RESPONSE: ", string(data))

}

/* ---------- POST Request with Form Data ----------- */
func postRequestWithFormData() {
	myUrl := "http://localhost:8080/postform"
	// Form Data format

	data := url.Values{}
	data.Add("name", "Elon Musk")
	data.Add("email", "elonmusk@gmail.com")
	data.Add("dob", "1971-06-28")
	data.Add("image", "https://upload.wikimedia.org/wikipedia/commons/8/85/Elon_Musk_Royal_Society_%28crop1%29.jpg")

	response, err := http.PostForm(myUrl, data)
	checkError(err)

	defer response.Body.Close()

	responseData, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Respones: ", string(responseData))
}

// functino for handling nil error
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
