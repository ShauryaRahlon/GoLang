package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Welcome to the files in Go lang")
	content := "This needs to go in a file"

	file, err := os.Create("./testingFile.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./testingFile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkError(err)
	fmt.Println("The data inside the file is \n", databyte)
}
