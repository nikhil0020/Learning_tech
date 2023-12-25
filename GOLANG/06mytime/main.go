package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()

	fmt.Println(presentTime)

	formatedTime := presentTime.Format("01-02-2006 15:04:05 Monday")

	// This is specific format "01-02-2006 15:04:05 Monday"

	fmt.Println(formatedTime)

	createdDate := time.Date(2020, time.April, 20, 12, 12, 50, 0, time.UTC)

	fmt.Println(createdDate)
}
