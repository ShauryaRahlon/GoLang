package main

import "fmt"

func main() {
	fmt.Println("welcome to arrays in golang")
	var list [4]string
	list[0] = "Apple"
	list[1] = "Mango"
	list[2] = "Banana"
	list[3] = "Grapes"
	fmt.Println("Fruit list is:", list)
	fmt.Println("length of array is :", len(list))

	var vegList = [3]string{"Potato", "Tomato", "Mushroom"}
	fmt.Println("Veg list is ", vegList)

}
