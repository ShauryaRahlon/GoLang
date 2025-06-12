package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our pizza")
	fmt.Println("please rate our pizza")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating our chicken", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //trims the useless space

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The converted number is", numRating+1)
	}
}
