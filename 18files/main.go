package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to File Handling")

	toBeWritten := "File handling in GO is cool."

	file, err := os.Create("./myfile.txt")

	// If file creating fails exit program with msg
	checkNilError(err)

	length, err := io.WriteString(file, toBeWritten)

	// If writing fails exit program with msg
	checkNilError(err)

	fmt.Println("Length is: ", length)

	defer file.Close()
	readFile("./myfile.txt")

}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)

	// If writing fails exit program with msg
	checkNilError(err)

	// filereading returns in byte format
	fmt.Println("Text inside file is: ", string(data))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
