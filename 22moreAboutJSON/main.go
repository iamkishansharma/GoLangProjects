package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/* If we don't start variable names with capital letters it will give nothing as output*/
type user struct {
	// mapping with api response keys
	Name  string `json:"name"`
	Email string `json:"email"`
	Dob   string `json:"dob"`
	Image string `json:"image,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")

	// Encode
	// EncodeJson()
	// Decode
	DecodeJson()

}

// encoding JSON
func EncodeJson() {
	myUsers := []user{
		{"Hitesh Choudhary", "hitesh@lco.dev", "1975-01-01", "https://avatars.githubusercontent.com/u/11613311?v=4"},
		{"Bryan Bolt", "bolt121@gmail.com", "1855-01-01", "https://wl-brightside.cf.tsp.li/resize/728x/jpg/978/58b/cecf2d560ca3833915b1fad736.jpg"},
		{"LCO", "lco123@gmail.com", "2015-12-12", ""},
	}

	// gives json unformatted data
	// jsonData, err := json.Marshal(myUsers)

	// gives json formatted data
	jsonData, err := json.MarshalIndent(myUsers, "", "  ")

	// checking for nil error
	checkError(err)

	fmt.Println("Json data is: ", string(jsonData))
}

// decoding JSON
func DecodeJson() {
	// getting json data from web
	// gowebserver should be running on localhost:8000
	response, err := http.Get("http://localhost:8000/get/users")
	checkError(err)
	fmt.Println("Response: ", response.StatusCode)

	byteJsonData, _ := ioutil.ReadAll(response.Body)

	// Will hold the JSON array comming from web as JSON format
	// USING user STRUCT
	var myUsers []user
	// is this data JSON?
	isValid := json.Valid(byteJsonData)

	if isValid {
		fmt.Println("----- Yay!, The data is valid JSON. ------")
		json.Unmarshal(byteJsonData, &myUsers)

		// fmt.Println(myUsers) // this will print all data without key
		fmt.Printf("%#v\n\n", myUsers)
	} else {
		fmt.Println("Not valid")
	}

	// NOT using user STRUCT instead using MAP array
	var myUsersMap []map[string]interface{}
	json.Unmarshal(byteJsonData, &myUsersMap)
	fmt.Printf("%#v\n", myUsersMap[0])

	for index := range myUsersMap {
		for k, v := range myUsersMap[index] {
			fmt.Printf("%v: %v Type: %T\n", k, v, v)
		}
	}

}

// handle error
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
