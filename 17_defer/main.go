package main

import "fmt"

func main() {

	fmt.Println("-------defer---------")

	defer fmt.Println("Hello World")
	defer fmt.Println("One")
	myDefer()
	defer fmt.Println("Two")

	fmt.Println("Khushi")
	fmt.Println("Khushi")
	fmt.Println("Khushi")

}

func myDefer()  {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i, "\n")
	}
}