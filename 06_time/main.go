package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to time study of golang.")

	presentTime := time.Now()
	// presentTimeNow := time.Now().Nanosecond()
	// fmt.Println(presentTimeNow)

	// this date is the standard for the formatting
	// given inside the documentation as well
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	fmt.Println(presentTime.Format("01-02-2006"))

	createdDate := time.Date(2020, time.March, 17, 23, 01, 0, 0, time.UTC)

	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}
