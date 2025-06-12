package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	langs := make(map[string]string)
	langs["JS"] = "javascript"
	langs["RB"] = "ruby"
	langs["SW"] = "swift"
	fmt.Println(langs)
	fmt.Println("JS is for", langs["JS"])

	delete(langs, "RB")
	fmt.Println(langs)

	//loops in golang
	for key, value := range langs {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
