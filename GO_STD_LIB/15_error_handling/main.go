package main

import "fmt"

func HandlingAnError() {
	categories := []string{"Watersports", "Chess", "Running"}
	for _, cat := range categories {
		total, err := Products.TotalPrice(cat)
		if err == nil {
			fmt.Println(cat, "Total:", ToCurrency(total))
		} else {
			fmt.Println(cat, "(no such category)")
		}
	}
}

func UsingAsyncMethod() {
	categories := []string{"Watersports", "Chess", "Running"}
	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
		} else {
			fmt.Println(message.Category, "(no such category)")
		}
	}
}

func TriggeringPanic() {
	categories := []string{"Watersports", "Chess", "Running"}
	channel := make(chan ChannelMessage, 10)
	// defer fmt.Println("Hello World")
	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
		} else {
			panic(message.CategoryError)
			//fmt.Println(message.Category, "(no such category)")
		}
	}
}

func recoveryFunc() {
	if arg := recover(); arg != nil {
		if err, ok := arg.(error); ok {
			fmt.Println("Error: ", err.Error())
		} else if str, ok := arg.(string); ok {
			fmt.Println("Message: ", str)
		} else {
			fmt.Println("Panic Recovered")
		}
	}
}

func RecoveringFromPanic() {
	defer recoveryFunc()
	categories := []string{"Watersports", "Chess", "Running"}
	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
		} else {
			panic(message.CategoryError)
			//fmt.Println(message.Category, "(no such category)")
		}
	}
}

type CategoryCountMessage struct {
	Category string
	Count    int
}

func processCategories(categories []string, outChan chan<- CategoryCountMessage) {
	defer func() {
		if arg := recover(); arg != nil {
			fmt.Println("RECOVERING IN GOROUTINE")
			fmt.Println(arg)
		}
	}()
	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			outChan <- CategoryCountMessage{
				Category: message.Category,
				Count:    int(message.Total),
			}
		} else {
			panic(message.CategoryError)
		}
	}
	close(outChan)
}

func RecoveringFromPanicInGoroutine() {
	categories := []string{"Watersports", "Chess", "Running"}
	channel := make(chan CategoryCountMessage)
	go processCategories(categories, channel)
	for message := range channel {
		fmt.Println(message.Category, "Total:", message.Count)
	}
}

func main() {
	// HandlingAnError()
	// UsingAsyncMethod()
	// TriggeringPanic()
	// RecoveringFromPanic()
	RecoveringFromPanicInGoroutine()
}
