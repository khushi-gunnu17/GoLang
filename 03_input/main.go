package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to this directory."
	fmt.Println(welcome)

	// walrus operator is used coz you don't know what will be the datatype.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rating for the pizza : ")

	// comma ok syntax || err ok syntax
	// we do not have try catch in golang
	// this syntax is like try, catch
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating : ", input)
	fmt.Printf("Type of this rating is : %T ", input)

}
