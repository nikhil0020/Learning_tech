package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	response, err := http.Get(url)

	checkNilErr(err)

	databytes, err := ioutil.ReadAll(response.Body)

	checkNilErr(err)

	defer response.Body.Close()
	content := string(databytes)

	fmt.Println(content)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
