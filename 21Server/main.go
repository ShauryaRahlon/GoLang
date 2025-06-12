package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("GET REQUESTS")
	// PerformGetRequests()
	// PerformPostRequests()
	FormOps()
}

func PerformGetRequests() {
	const myUrl = "http://localhost:8000/get"

	res, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("Status: ", res.StatusCode)
	fmt.Println("Content Length: ", res.ContentLength)

	// content, _ := ioutil.ReadAll(res.Body)  this is deprecated
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))

}

func PerformPostRequests() {
	const myUrl = "http://localhost:8000/post"
	reqBody := strings.NewReader(`
		{
			"coursename":"Solving Cpp"
		}
	`)

	res, err := http.Post(myUrl, "application/json", reqBody)

	if err != nil {
		panic((err))
	}
	defer res.Body.Close()
	cont, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cont))
}

func FormOps() {
	const myUrl = "http://localhost:8000/postform"

	//form data

	data := url.Values{}

	data.Add("firstname", "shaurya")
	data.Add("lastname", "rahlon")

	res, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	cont, _ := io.ReadAll(res.Body)
	fmt.Println(string(cont))
}
