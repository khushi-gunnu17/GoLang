package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	// "time"
	"crypto/rand"
)

const CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) (string, error)  {

	b := make([]byte, length)

	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(CHARSET))))

		if err != nil {
			return "", err
		}

		b[i] = CHARSET[num.Int64()]
	}
	return string(b), nil
}


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

	length := 10
	randomStr, err := randomString(length)

	if err != nil {
		fmt.Println("Error generating random string:", err)
		return
	}

	fmt.Println("Random String : ", randomStr)

}
