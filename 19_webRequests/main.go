package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "https://google.com"

func main() {

	fmt.Println("-------------------Handling web requests-------------------")

	response, err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Reponse is of type: %T\n", response) // referennce of the original response

	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)

}
