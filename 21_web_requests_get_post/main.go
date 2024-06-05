package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println("-----------web requests------------")

	// PerformGetRequest()

	// PerformPostJSONRequest()

	PerformPostFormRequests()

}

func PerformGetRequest()  {

	const MYURL = "http://localhost:8000/get"

	response, err := http.Get(MYURL)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Content Length is : ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("")
	// fmt.Println(string(content))			// without a package
	fmt.Println("Byte count is : ", byteCount)
	fmt.Println(responseString.String())

	fmt.Println("")
	fmt.Printf("responseString: %v\n", responseString)

}


func PerformPostJSONRequest()  {

	const MYURL = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename" : "Go Language",
			"Price" : 1000,
			"platform" : "KhushiLearnWebiste.in"
		}
	`)

	response, err := http.Post(MYURL, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}


func PerformPostFormRequests()  {
	
	const MYURL = "http://localhost:8000/postform"

	// formdata
	data := url.Values{}
	data.Add("firstname", "Khushi")
	data.Add("lastname", "Sharma")
	data.Add("email", "Khushi@gmail.com")

	response, err := http.PostForm(MYURL, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)		

	fmt.Println(string(content))

}