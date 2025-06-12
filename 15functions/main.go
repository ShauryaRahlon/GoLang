package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Welcome to functions in go lang")
	res := adder(3, 4)
	fmt.Println("Result is:", res)

	proRes := proAdder(1, 3, 5, 1)
	fmt.Println(proRes)
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func adder(valOne int, valTwo int) int /* add return value data type */ {
	return valOne + valTwo
}
func greeter() {
	fmt.Println("Hi from Go lang")
}
