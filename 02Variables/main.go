package main

import "fmt"

const LoginToken string = "asdadas"

func main() {
	var username string = "shaurya"
	fmt.Println(username)
	fmt.Printf("variable is of :%T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of :%T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Print(smallVal)

	var smallFloat float32 = 255.434334123132
	fmt.Print(smallFloat)
	fmt.Printf("variable type :%T \n", smallFloat)

	var anotherVar int
	fmt.Println(anotherVar) //doesnt put garbage value in it
	fmt.Printf(" this is %T\n", anotherVar)

	//implicit type
	var web = "solve.com"
	fmt.Println(web)

	numberOfUser := 30000
	fmt.Println(numberOfUser)
	fmt.Println(LoginToken)

}
