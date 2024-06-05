package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("------switch------")
	
	rand.Seed(time.Now().UnixNano())		
	diceNum := rand.Intn(6) + 1

	fmt.Println("Value of dice is : ", diceNum)

	switch diceNum {
	case 1 : 
		fmt.Println("Dice value is 1 and you can start your game.")
		
	case 2 : 
		fmt.Println("Dice value is 2 and you can move 2 spaces.")
		
	case 3 : 
		fmt.Println("Dice value is 3 and you can move 3 spaces.")
		// fallthrough	
		
	case 4 : 
		fmt.Println("Dice value is 4 and you can move 4 spaces.")
		// fallthrough

	case 5 : 
		fmt.Println("Dice value is 5 and you can move 5 spaces.")
		
	case 6 : 
		fmt.Println("Dice value is 6 and you can move 6 spaces and roll the dice again")
		
	default : 
		fmt.Println("No idea!")

	}

}
