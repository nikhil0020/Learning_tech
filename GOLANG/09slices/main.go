package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Mango", "Apple", "Orange", "Banana"}

	fmt.Println(fruitList)
	// Slices are most used data structure in GOLANG

	fruitList = append(fruitList, "PineApple", "Grapes")

	fmt.Println(fruitList)
	fruitList = append(fruitList[1:4])
	fmt.Println(fruitList)

	// * Another way to declare slices

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	// highScores[4] = 234
	// above line will give errors , runtime error: index out of range [4] with length 4

	highScores = append(highScores, 912)
	// above line will not give error

	sort.Ints(highScores)
	fmt.Println(highScores)

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)
}
