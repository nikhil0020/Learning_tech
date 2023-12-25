package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Welcome to handle url")
	const URL string = "https://lca.dev:3000/learn?coursename=reactjs&paymentid=12124fsafsdff"

	response, err := url.Parse(URL)

	checkNilErr(err)

	fmt.Println(response.Host)
	fmt.Println(response.Path)
	fmt.Println(response.RawQuery)

	qparams := response.Query()
	fmt.Println(qparams)
	fmt.Println(qparams["coursename"])

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
