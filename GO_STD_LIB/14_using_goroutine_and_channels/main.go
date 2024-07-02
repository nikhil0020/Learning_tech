package main

import (
	"fmt"
	"math/rand"
	"time"
)

func WhenNumberOfValuesAreKnown() {
	CalcStoreTotal(Products)
}

func UsingForLoopForUnknownAmountOfValues() {
	dispatchChannel := make(chan DispatchNotification, 100)
	go DispatchOrders(dispatchChannel)
	for {
		details := <-dispatchChannel
		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
			"x", details.Product.Name)
	}
	// This will give deadlock error
	// because if there is no value then also for will be looking for values.
}

func CheckingForClosedChannel() {
	dispatchChannel := make(chan DispatchNotification, 100)
	go DispatchOrders(dispatchChannel)
	for {
		if details, open := <-dispatchChannel; open {
			fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
		} else {
			fmt.Println("Channel has been closed")
			break
		}
	}
}

func EnumeratingTheChannel() {
	dispatchChannel := make(chan DispatchNotification, 100)
	go DispatchOrders(dispatchChannel)
	for details := range dispatchChannel {
		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
			"x", details.Product.Name)
	}
	fmt.Println("Channel has been closed")
}

// Restricting the direction of channel

func MistakingOperations() {
	channel := make(chan DispatchNotification, 100)
	rand.Seed(time.Now().UTC().UnixNano())
	orderCount := rand.Intn(3) + 2
	fmt.Println("Order count:", orderCount)
	for i := 0; i < orderCount; i++ {
		channel <- DispatchNotification{
			Customer: Customers[rand.Intn(len(Customers)-1)],
			Quantity: rand.Intn(10),
			Product:  ProductList[rand.Intn(len(ProductList)-1)],
		}
		if i == 1 {
			notification := <-channel
			fmt.Println("Read:", notification.Customer)
		}
	}
	close(channel)
}

// Restricing Channel Argument Direction

func receiveDispatches(channel <-chan DispatchNotification) {
	for details := range channel {
		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
			"x", details.Product.Name)
	}
	fmt.Println("Channel has been closed")
}

func RestricingChannel() {
	dispatchChannel := make(chan DispatchNotification, 100)
	var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
	var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel
	go DispatchOrdersRestricted(sendOnlyChannel)
	receiveDispatches(receiveOnlyChannel)
}

func UsingExplicitConversionForChannels() {
	dispatchChannel := make(chan DispatchNotification, 100)

	go DispatchOrdersRestricted(chan<- DispatchNotification(dispatchChannel))
	receiveDispatches((<-chan DispatchNotification)(dispatchChannel))
}

// Using a select statement

func UsingSelectStatement() {
	dispatchChannel := make(chan DispatchNotification, 100)
	go DispatchOrdersSelect(chan<- DispatchNotification(dispatchChannel))
	// receiveDispatches((<-chan DispatchNotification)(dispatchChannel))
	for {
		select {
		case details, ok := <-dispatchChannel:
			if ok {
				fmt.Println("Dispatch to", details.Customer, ":",
					details.Quantity, "x", details.Product.Name)
			} else {
				fmt.Println("Channel has been closed")
				goto alldone
			}
		default:
			fmt.Println("-- No message ready to be received")
			time.Sleep(time.Millisecond * 500)
		}
	}
alldone:
	fmt.Println("All values received")
}

// Receiving from Multiple channel

func enumerateProducts(channel chan<- *Product) {
	for _, p := range ProductList[:3] {
		channel <- p
		time.Sleep(time.Millisecond * 800)
	}
	close(channel)
}

func ReceivingFromMultipleChannel() {
	dispatchChannel := make(chan DispatchNotification, 100)
	go DispatchOrdersSelect(chan<- DispatchNotification(dispatchChannel))
	productChannel := make(chan *Product)
	go enumerateProducts(productChannel)
	openChannels := 2
	for {
		select {
		case details, ok := <-dispatchChannel:
			if ok {
				fmt.Println("Dispatch to", details.Customer, ":",
					details.Quantity, "x", details.Product.Name)
			} else {
				fmt.Println("Dispatch channel has been closed")
				dispatchChannel = nil
				openChannels--
			}
		case product, ok := <-productChannel:
			if ok {
				fmt.Println("Product:", product.Name)
			} else {
				fmt.Println("Product channel has been closed")
				productChannel = nil
				openChannels--
			}
		default:
			if openChannels == 0 {
				goto alldone
			}
			fmt.Println("-- No message ready to be received")
			time.Sleep(time.Millisecond * 500)
		}
	}

alldone:
	fmt.Println("All values received")
}

func enumerateProductsForSending(channel chan<- *Product) {
	for _, p := range ProductList {
		select {
		case channel <- p:
			fmt.Println("Sent product:", p.Name)
		default:
			fmt.Println("Discarding product:", p.Name)
			time.Sleep(time.Second)
		}
	}
	close(channel)
}

func SendingUsingSelect() {
	productChannel := make(chan *Product, 5)
	go enumerateProductsForSending(productChannel)
	time.Sleep(time.Second)
	for p := range productChannel {
		fmt.Println("Received product:", p.Name)
	}
}

func enumerateProductsForMultipleChannel(channel1, channel2 chan<- *Product) {
	for _, p := range ProductList {
		select {
		case channel1 <- p:
			fmt.Println("Send via channel 1")
		case channel2 <- p:
			fmt.Println("Send via channel 2")
		}
	}
	close(channel1)
	close(channel2)
}

func SendingOverMultipleChannel() {
	c1 := make(chan *Product, 2)
	c2 := make(chan *Product, 2)
	go enumerateProductsForMultipleChannel(c1, c2)
	time.Sleep(time.Second)
	for p := range c1 {
		fmt.Println("Channel 1 received product:", p.Name)
	}
	for p := range c2 {
		fmt.Println("Channel 2 received product:", p.Name)
	}
}

func main() {
	fmt.Println("main function started")
	// WhenNumberOfValuesAreKnown()
	// UsingForLoopForUnknownAmountOfValues()
	// CheckingForClosedChannel()
	// EnumeratingTheChannel()
	// MistakingOperations()
	// RestricingChannel()
	// UsingExplicitConversionForChannels()
	// UsingSelectStatement()
	// ReceivingFromMultipleChannel()
	// SendingUsingSelect()
	SendingOverMultipleChannel()
	// time.Sleep(time.Second * 5)
	fmt.Println("main function ended")
}
