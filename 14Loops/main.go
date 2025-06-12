package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")
	days := []string{"sunday", "monday", "thursday", "saturday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for _, day := range days {
		fmt.Printf("index is and the value is %v\n", day)
	}

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 4 {
			break
		}

		fmt.Println("Value is:", rougueValue)
		rougueValue++
	}
}
