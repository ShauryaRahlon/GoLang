package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//no inheritance in go lang no super or parent

	shaurya := User{"Shaurya", "shaurya@mail.in", true, 32}
	fmt.Println(shaurya)
	fmt.Printf("details are: %+v\n", shaurya)
	fmt.Printf("Email: %+v\n", shaurya.Email)
	shaurya.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User Active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
