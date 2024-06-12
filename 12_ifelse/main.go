package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("-------if-else--------")

	fmt.Print("Enter the login count : ")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	loginCount, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)


	var result string

	if loginCount > 10 {
		result = "Regular user"
	} else if loginCount < 5 {
		result = "Not a regular user"
	} else {
		result = "Normal user"
	}

	fmt.Println(result)



	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd.")
	}



	// using a variable while declaring it, common syntax in golang.
	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is greater than 10")
	}



	// if err != nil {}

}
