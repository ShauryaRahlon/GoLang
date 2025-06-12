package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// rand.Seed(time.Now().Unix())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1")
		fallthrough
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("what was that")
	}

}
