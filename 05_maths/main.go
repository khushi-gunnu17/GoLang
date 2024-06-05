package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	// "time"
	"crypto/rand"
)

func main() {

	fmt.Println("---------maths----------")

	var myNumOne int = 2
	var myNumTwo float64 = 4.5

	fmt.Println("The sum is : ", myNumOne + int(myNumTwo))


	// random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)


	// random from crpto 
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(100))
	fmt.Println(myRandomNum)

}
