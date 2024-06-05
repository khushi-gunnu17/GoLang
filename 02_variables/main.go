package main

// formatted I/O package = import "fmt"
import "fmt"

// if first char in the variable is capital, so then it is made a public variable.
const LoginToken string = "ndhbxusbjnaj"

// walrus operator not allowed in global but inside the method can be used
// jwtToken := 300000


func main()  {
	
	var username string = "Khushi Sharma"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)
	fmt.Println("")

	
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)
	fmt.Println("")
	

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)
	fmt.Println("")
	

	var smallFloat float32 = 255.25536263762182
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)
	fmt.Println("")
	
	
	// default values and some aliases
	var anotherVariable int		// defaults to 0
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)
	fmt.Println("")
	
	
	// implicit way of declaring the variables
	var website = "google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type : %T \n", website)
	fmt.Println("")
	
	
	// no var style
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)
	fmt.Println("")

	
	// const variable
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)

}