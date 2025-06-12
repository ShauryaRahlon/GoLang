package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")
	var FruitList = []string{"Apple"}
	fmt.Printf("type of list is %T\n", FruitList)

	FruitList = append(FruitList, "Mangos", "Apple")
	// fmt.Println(FruitList)

	FruitList = append(FruitList[:2])
	// fmt.Println(FruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 240
	highScores[2] = 900
	highScores[3] = 100

	highScores = append(highScores, 332, 22, 1102)
	// fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)

	// fmt.Println(highScores)

	//how to remove value from slices based on their index

	var course = []string{"react", "node", "swift", "ruby"}
	fmt.Println(course)

	var index int = 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)

}
