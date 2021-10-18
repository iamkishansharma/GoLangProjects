package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iamkishansharma/GoLangProjects/tree/25ApiWithMongoDB/router"
)

func main() {
	fmt.Println("Welcome to Go Api with MongoDB")
	r := router.Router()
	fmt.Println("Server started...")

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listining at port 4000 ...")

}
