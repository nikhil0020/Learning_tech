package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Printf("Welcome to file handling\n")

	content := "This is the content of the file"

	// Creating the file

	file, err := os.Create("myfile.txt")

	checkNilErr(err)

	// * Writing in file

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Printf("Length is : %v\n", length)

	readFile("./myfile.txt")
}

func readFile(filename string) {

	// Reading data from file in bytes

	databytes, err := ioutil.ReadFile(filename)

	checkNilErr(err)

	// Databytes are in byte form so we have to convert it to string

	content := string(databytes)

	fmt.Printf("Content of the file is : %v\n", content)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
