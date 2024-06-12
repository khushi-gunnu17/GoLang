package main

// In the root directory, there should be only one file.

import (
	"fmt"
	"log"
	"mongoapi/router"
	"net/http"
)

func main() {

	fmt.Println("--------------------mongoDb-api-------------------")

	// calling the router
	r := router.Router()

	fmt.Println("Server is getting started......")

	log.Fatal(http.ListenAndServe(":6000", r))

	fmt.Println("Listening at port 6000.....")

}
