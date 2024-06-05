package main

import "fmt"

func main() {

	fmt.Println("---------methods--------")

	khushi := User{"Khushi", "Khushi@google.com", true, 20}
	fmt.Println(khushi)

	fmt.Printf("Khushi details are : %+v\n", khushi)
	fmt.Printf("Name is : %v and Email is : %v\n", khushi.Name, khushi.Email)
	
	
	khushi.GetStatus()

	khushi.NewMail()

	// didn't change the email address in actuality
	fmt.Println(khushi.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// if planning to exporth this methid, syntax changes
func (u User) GetStatus() {
	fmt.Println("Is user active? : ", u.Status)
}

// They passes copies of the values
func (u User) NewMail()  {
	u.Email = "test@gmail.com"
	fmt.Println("Email of this user is : ", u.Email)
}