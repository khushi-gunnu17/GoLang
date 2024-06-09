package main

import "fmt"

func main() {

	fmt.Println("Welcome to maps in golang")
	fmt.Println("")

	// key : value
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["C++"] = "C++"

	fmt.Println("List of all languages : ", languages)
	fmt.Println("")
	fmt.Println("JS shorts for : ", languages["JS"])
	fmt.Println("")
	
	// deleting some values
	delete(languages, "C++")
	fmt.Println("List of all languages : ", languages)
	fmt.Println("")

	// looping in maps
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
	fmt.Println("")

	for _, value := range languages {
		fmt.Printf("The values are : %v\n", value)
	}

}
