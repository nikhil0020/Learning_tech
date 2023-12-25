package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Request")
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	checkNilErr(err)

	defer response.Body.Close()

	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Content length is : ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)

	fmt.Println("byteCount is : ", byteCount)

	fmt.Println("resposeString is : ", responseString.String())
	// fmt.Println(string(content))
}

func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"Brand":"OneFourth",
			"platform" : "Online"
		}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)

	checkNilErr(err)

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	var responseString strings.Builder

	responseString.Write(content)
	fmt.Println("Response string is : ", responseString.String())
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// formdata

	data := url.Values{}

	data.Add("firstname", "Nikhil")
	data.Add("lastname", "Singh")
	data.Add("email", "nikhil@go.dev")

	response, err := http.PostForm(myurl, data)

	checkNilErr(err)

	defer response.Body.Close()

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)

	responseString.Write(content)

	fmt.Println(responseString.String())
}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
