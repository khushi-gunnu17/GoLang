package main

import "fmt"

func main() {

	fmt.Println("-------functions--------")	

	greeter()


	// you are not allowed to write functions inside the function
	// func greeterTwo()  {
	// 	fmt.Println("Another Hello")
	// }

	greeterTwo()


	result := addTwo(3, 4)
	fmt.Println("Result is : ", result)

	proResult, myMessage := proAdder(2, 4, 6, 8, 10)
	fmt.Println("The result of many numbers addition is : ", proResult)
	fmt.Println("The message of many numbers addition is : ", myMessage)

}

func greeter() {
	fmt.Println("Namaste from golang")
}

func greeterTwo()  {
	fmt.Println("Another Hello")
}

func addTwo(valOne int, valTwo int) int  {		// function signature, which is the datatype of the return value
	return valOne + valTwo
}

func proAdder(values ...int) (int, string)  {
	total := 0

	for _ , val := range values {
		total += val
	}

	return total, "Hi pro result func"
}