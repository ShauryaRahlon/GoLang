package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("world 2")
	defer fmt.Println("world 3")
	fmt.Println("Hello ")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
