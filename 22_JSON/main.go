package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string	`json:"coursename"`
	Price    int	
	Platform string	`json:"website"`
	Password string	`json:"-"`
	Tags     []string	`json:"tags,omitempty"`
}

func main() {

	fmt.Println("------------------JSON--------------------")

	// encoding of the json
	// EncodeJson()

	// decoding of the json
	DecodeJson()
}

func EncodeJson() {

	myCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "bdyeguv78", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 200, "LearnCodeOnline.in", "bhdva", []string{"full-stack", "js"}},
		{"ANgular Bootcamp", 299, "LearnCodeOnline.in", "huceiua", nil},
	}

	// package this data as JSON data 

	finaljson, err := json.MarshalIndent(myCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finaljson)

}


func DecodeJson()  {
	
	jsonDataFromWeb := []byte(`
		{
			"coursename": "ReactJS Bootcamp",
			"Price": 299,
			"website": "LearnCodeOnline.in",
			"tags": ["web-dev", "js"]
		}
	`)

	var myCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid.")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID.")
	}
	fmt.Println("")



	// some cases where you just want to add data to key value pair

	// if you don't know what type of data will be coming into the map, you can add interface to it.
	var myOnlineData map[string]interface{}
	
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)
	fmt.Println("")

	for key, value := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is : %T\n", key, value, value)
	}
	fmt.Println("")

}
