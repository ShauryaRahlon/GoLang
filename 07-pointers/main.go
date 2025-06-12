package main

import "fmt"

func main() {
	fmt.Println("Welcome")

	var ptr *int
	fmt.Println(ptr)

	myNumber := 23
	var ptr1 = &myNumber
	fmt.Println("Value of actual pointer is ", ptr1)
	fmt.Println("Value of actual pointer is ", *ptr1)
}
