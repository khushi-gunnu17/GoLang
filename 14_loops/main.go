package main

import "fmt"

func main() {

	fmt.Println("------loops-------")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	fmt.Println(days)
	fmt.Println("")
	

	// type 1
	// for i := 0; i<len(days); i++ {
	// 	fmt.Println(days[i])
	// }


	// type 2
	// i gives the index value
	// for i := range days {
	// 	fmt.Println(days[i])
	// }


	// type 3
	for index, day := range days {
		fmt.Printf("Index is %v and the day is %v\n", index, day)
	}
	fmt.Println("")



	rogueValue := 1

	for rogueValue < 11 {

		// break
		// if rogueValue == 5 {
		// 	break
		// }

		// continue
		if rogueValue == 5 {
			rogueValue++
			continue
		}

		// goto
		if rogueValue == 8 {
			goto location
		}

		fmt.Println("Value is :", rogueValue)
		rogueValue++
	}

location :
	fmt.Println("Jumping at google.com")

	fmt.Println("")



	for i:=0; i < 5; i++ {
		if i == 3 {
		  continue
		}
		fmt.Println(i)
	}


}
