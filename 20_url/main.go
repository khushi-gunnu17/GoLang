package main

import (
	"fmt"
	"net/url"
)

const URL string = "https://lco.dev:4000/learn?coursename=reactjs&paymentid=njbddbdhebjdjjsn"

func main() {

	fmt.Println("------------------urls-------------------")

	fmt.Println(URL)
	fmt.Println("")

	// parsing the URL

	result, _ := url.Parse(URL)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	fmt.Println("")

	qparams := result.Query()

	fmt.Printf("The type of qparams are: %T\n", qparams)
	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Param is : ", val)
	}
	fmt.Println("")




	// to construct a url from scratch
	// you want to pass the reference and not the copy
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=khushi",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
	fmt.Println("")

}
