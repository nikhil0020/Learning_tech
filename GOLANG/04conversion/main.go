package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Please rate between 1 and 5 : ")
	input, _ := reader.ReadString('\n')

	fmt.Printf("Thanks for rating %s", input)

	numRating, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	numRating++
	fmt.Println(numRating)

}
