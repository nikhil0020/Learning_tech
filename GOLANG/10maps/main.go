package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	// 	var a = make(map[KeyType]ValueType)
	// b := make(map[KeyType]ValueType)

	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	// used to delete a value from a map
	fmt.Printf("a\t%v\n", a)
	fmt.Println(b)
	fmt.Println(languages)

	delete(languages, "RB")
	// Loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("For key %v , value is %v\n", key, value)
	}

}
