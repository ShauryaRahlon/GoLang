package main

import "fmt"

func main() {
	fmt.Println("If else in go lang")
	loginCnt := 23
	var res string
	if loginCnt < 10 {
		res = "regular user"
	} else {
		res = "abnormal user"
	}
	fmt.Println("res is :", res)

	if 8%2 == 0 {
		fmt.Println("Even Num")
	} else {
		fmt.Println("Odd")
	}

	if num := 3; num < 10 { //we assigned it here and checked it here in place useful in checking web requests and api
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

}
