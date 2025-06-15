package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string
	Password string   `json:"-"` //this wont be reflected
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON details")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourse := []course{
		{"React Bootcamp", 400, "solve.com", "abc123", []string{"web-dev", "js"}},
		{"GO bootcamp", 400, "solve.com", "abc123", []string{"backend-dev", "js"}},
		{"Angular Bootcamp", 400, "solve.com", "abc123", []string{"Frontend-dev", "js"}},
		{"Debate Bootcamp", 400, "solve.com", "Shaurya", nil},
	}

	//package this as json
	// finalJson, err := json.Marshal(lcoCourse)
	finalJson, err := json.MarshalIndent(lcoCourse, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`
		 {
                "courseName": "Angular Bootcamp",
                "Price": 400,
                "Platform": "solve.com",
                "tags": ["Frontend-dev", "js"]
        }
	`)

	var checkdata course

	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonData, &checkdata) //check data is passed as a reference
		fmt.Printf("%#v\n", checkdata)
	} else {
		fmt.Println("Json was not valid")
	}

	//some cases where u just want to add data to key value

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Printf("%#v\n", checkdata)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is : %T\n", k, v, v)
	}

}
