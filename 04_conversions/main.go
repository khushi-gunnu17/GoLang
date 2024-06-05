package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our pizza app.")

	fmt.Println("Please rate our pizza between 1 and 5 : ")

	reader := bufio.NewReader(os.Stdin)

	// comma, ok syntax
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating : ", input)

	// conversion here through the package strconv
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
		// panic(err)		// it will cause the program to end here
	} else {
		fmt.Println("Adding 1 to you rating : ", numRating + 1)
	}

}
