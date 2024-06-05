package main

import "fmt"

func main() {

	fmt.Println("Welcome to array in GoLang.")	

	var fruitList [4]string
	fruitList[0] = "Banana"
	fruitList[1] = "Apple"
	fruitList[3] = "Kiwi"

	fmt.Println("Fruit list is : ", fruitList)
	fmt.Println("Fruit list is : ", len(fruitList))


	
	var vegList = [5]string{"Potato", "Cauliflower", "Mushroom"}

	fmt.Println("Veg List is : ", vegList)
	fmt.Println("Veg List is : ", len(vegList))

}