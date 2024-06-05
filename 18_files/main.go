package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("------------Managing-files--------------")

	content := "This needs to go in a file = google.com"

	file, err := os.Create("./18_files/myLocationFile.txt")

	if err != nil {
		panic(err)		// shuts down the execution of the program
	}

	// writing into the file
	length, err := io.WriteString(file, content)

	checkNilError(err)
	fmt.Println("Length is : ", length)

	defer file.Close()


	// reading the file
	readFile("./18_files/myLocationFile.txt")

}


func readFile(filename string)  {

	databyte , err := os.ReadFile(filename)

	checkNilError(err)

	fmt.Println("Text data inside the file is \n", string(databyte))

}


func checkNilError(err error)  {
	if err != nil {
		panic(err)
	}
}