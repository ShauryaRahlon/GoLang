package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//no inheritance in go lang no super or parent

	shaurya := User{"Shaurya", "shaurya@mail.in", true, 32}
	fmt.Println(shaurya)
	fmt.Printf("details are: %+v\n", shaurya)
	fmt.Printf("Email: %+v\n", shaurya.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
