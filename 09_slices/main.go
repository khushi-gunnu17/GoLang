package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to the topic slices.")

	// syntax for slices
	var fruitList = []string{"Apple", "Banana", "Kiwi"}
	fmt.Printf("Type of fruitlist is : %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Peach")
	fmt.Println(fruitList)

	// to make slices
	// fruitList = append(fruitList[1 : 3])
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[: 4])
	// fmt.Println(fruitList)
	fmt.Println("")






	highScores := make([]int, 4)

	// default allocation of the memory
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777

	// reallocation of the memory and all the values will be accommodated
	highScores = append(highScores, 555, 444)

	fmt.Println(highScores)
	fmt.Println("")





	// sorting
	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	fmt.Println("")





	// to remove a value from a slice based on a index
	var courses = []string{"Python", "React.js", "Swift", "Ruby", "JavaScript"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index + 1 : ]...)
	fmt.Println(courses)
	fmt.Println("")







	// copy function
	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))
	fmt.Println("")



	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	fmt.Println(neededNumbers)
	fmt.Printf("length = %d\n", len(neededNumbers))
	fmt.Printf("capacity = %d\n", cap(neededNumbers))
	fmt.Println("")
	
	// creating copy
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))

}
