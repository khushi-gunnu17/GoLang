package main

import "fmt"

func main() {

	fmt.Println("-----Structs------")

	// no inheritance, super keyword, or the concept of parent class in golang

	khushi := User{"Khushi", "Khushi@google.com", true, 20}
	fmt.Println(khushi)

	fmt.Printf("Khushi details are : %+v\n", khushi)
	fmt.Printf("Name is : %v and Email is : %v\n", khushi.Name, khushi.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
