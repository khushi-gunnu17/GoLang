package main

import "fmt"

func main() {

	fmt.Println("Welcome to class on Pointers")

	// var ptr *int 
	// fmt.Println("Address of pointer is : ", ptr)			// default <nil>

	myNum := 23

	var ptr1 = &myNum		// reference to a memory
	fmt.Println("Address of pointer is : ", ptr1)
	fmt.Println("Value of pointer is : ", *ptr1)

	*ptr1 = *ptr1 * 2
	fmt.Println("New Value after mul is : ", myNum)

}
