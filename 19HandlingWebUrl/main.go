package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.prik.dev/"

func main() {
	fmt.Println("Web Requests")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type %T\n", res)
	defer res.Body.Close()

	dataByte, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(dataByte)
	fmt.Println(content)
}
