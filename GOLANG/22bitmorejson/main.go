package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	// EncodeJson()
	DecodedJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"Reactjs", 99, "onefourth.com", "abc123", []string{"webdev", "js"}},
		{"Mern", 199, "onefourth.com", "abcd123", []string{"mern", "js"}},
		{"Angular", 299, "onefourth.com", "abcde123", nil},
	}

	// finaljson, err := json.Marshal(lcoCourses)
	finaljson, err := json.MarshalIndent(lcoCourses, "", "\t")

	checkNilErr(err)

	fmt.Printf("%s\n", finaljson)
}

func DecodedJson() {
	jsonDataFromWeb := []byte(`
		{
			"coursename": "Reactjs",
			"Price": 99,
			"website": "onefourth.com",
			"tags": [
					"webdev",
					"js"
			]
		}
	`)

	checkValid := json.Valid(jsonDataFromWeb)

	var lcoCourse course
	if checkValid {
		fmt.Println("JSON data is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON data is invalid")
	}

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v Value is %v and Type is %T\n", k, v, v)
	}
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
