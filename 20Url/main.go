package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://docs.google.com/forms/d/e/1FAIpQLSfSFwlpjG-4wJvhhoUVTlEBxZ86Dk_MAr8vVS3vVioE6NMTHw/viewform"

func main() {
	fmt.Println("Welcome to handling urls")

	res, _ := url.Parse(myUrl)
	// fmt.Println(res.Scheme)
	// fmt.Println(res.Host)
	// fmt.Println(res.Port())
	fmt.Println(res.Query())
	params := res.Query()
	fmt.Printf("This type of query are: %T\n", params)

}
