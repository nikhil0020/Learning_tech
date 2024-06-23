package main

import (
	"fmt"
	"time"
)

func PrintTime(label string, t *time.Time) {
	Printfln("%s: Day: %v: Month: %v Year: %v",
		label, t.Day(), t.Month(), t.Year())
}

func PrintTimeLayout(label string, t *time.Time) {
	layout := "Day: 02 Month: Jan Year: 2006"
	fmt.Println(label, t.Format(layout))
}

func PrintTimeDefinedLayout(label string, t *time.Time) {
	//layout := "Day: 02 Month: Jan Year: 2006"
	fmt.Println(label, t.Format(time.RFC822Z))
}

func CreatingTimeValues() {
	current := time.Now()
	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(1433228090, 0)
	PrintTime("Current", &current)
	PrintTime("Specific", &specific)
	PrintTime("UNIX", &unix)
}

func CreateFormattedTimeValues() {
	current := time.Now()
	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(1433228090, 0)
	PrintTimeLayout("Current", &current)
	PrintTimeLayout("Specific", &specific)
	PrintTimeLayout("UNIX", &unix)
}

func CreatingTimeWithDefinedFormat() {
	current := time.Now()
	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(1444338090, 0)
	PrintTimeDefinedLayout("Current", &current)
	PrintTimeDefinedLayout("Specific", &specific)
	PrintTimeDefinedLayout("UNIX", &unix)
}

func ParsingTimeValueWithString() {
	layout := "2006-Jan-02"

	dates := []string{
		"1995-Jun-09",
		"2015-Jun-02",
	}

	for _, d := range dates {
		time, err := time.Parse(layout, d)

		if err == nil {
			PrintTimeDefinedLayout("Parsed: ", &time)
		} else {
			Printfln("Error %s", err.Error())
		}
	}
}

func UsingPredefinedDateLayout() {
	dates := []string{
		"09 Jun 95 00:00 GMT",
		"02 Jun 15 00:00 GMT",
	}

	for _, d := range dates {
		time, err := time.Parse(time.RFC822, d)
		if err == nil {
			PrintTimeDefinedLayout("Parsed", &time)
		} else {
			Printfln("Error: %s", err.Error())
		}
	}
}

func SpecifyingParsingLocation() {
	layout := "02 Jan 06 15:04"
	date := "09 Jun 95 19:30"
	london, lonerr := time.LoadLocation("Europe/London")
	newyork, nycerr := time.LoadLocation("America/New_York")
	if lonerr == nil && nycerr == nil {
		nolocation, _ := time.Parse(layout, date)
		londonTime, _ := time.ParseInLocation(layout, date, london)
		newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
		PrintTimeDefinedLayout("No location:", &nolocation)
		PrintTimeDefinedLayout("London:", &londonTime)
		PrintTimeDefinedLayout("New York:", &newyorkTime)
	} else {
		fmt.Println(lonerr.Error(), nycerr.Error())
	}
}

// Using the local location

// If the place name used to create a Location is Local,
// then the time zone setting of the machine running the application is used,

func UsingLocalLocation() {
	layout := "02 Jan 06 15:04"
	date := "09 Jun 95 19:30"

	london, lonerr := time.LoadLocation("Europe/London")
	newyork, nyerr := time.LoadLocation("America/New_York")

	localTime, _ := time.LoadLocation("Local")

	if lonerr == nil && nyerr == nil {
		nolocation, _ := time.Parse(layout, date)

		londonTime, _ := time.ParseInLocation(layout, date, london)
		newyorkTime, _ := time.ParseInLocation(layout, date, newyork)

		localTime, _ := time.ParseInLocation(layout, date, localTime)

		PrintTimeDefinedLayout("No location:", &nolocation)
		PrintTimeDefinedLayout("London:", &londonTime)
		PrintTimeDefinedLayout("New York:", &newyorkTime)
		PrintTimeDefinedLayout("Local:", &localTime)
	}

}

// Specifying Time Zones Directly

func SpecifyingTimeZoneDirectly() {
	layout := "02 Jan 06 15:04"
	date := "09 Jun 95 19:30"
	london := time.FixedZone("BST", 1*60*60)
	newyork := time.FixedZone("EDT", -4*60*60)
	local := time.FixedZone("Local", 0)
	nolocation, _ := time.Parse(layout, date)
	londonTime, _ := time.ParseInLocation(layout, date, london)
	newyorkTime, _ := time.ParseInLocation(layout, date, newyork)
	localTime, _ := time.ParseInLocation(layout, date, local)
	PrintTimeDefinedLayout("No location:", &nolocation)
	PrintTimeDefinedLayout("London:", &londonTime)
	PrintTimeDefinedLayout("New York:", &newyorkTime)
	PrintTimeDefinedLayout("Local:", &localTime)
}

// Representing duration

func DurationMethods() {
	var d time.Duration = time.Hour + (30 * time.Minute)
	Printfln("Hours: %v", d.Hours())
	Printfln("Mins: %v", d.Minutes())
	Printfln("Seconds: %v", d.Seconds())
	Printfln("Millseconds: %v", d.Milliseconds())
	rounded := d.Round(time.Hour)
	Printfln("Rounded Hours: %v", rounded.Hours())
	Printfln("Rounded Mins: %v", rounded.Minutes())
	trunc := d.Truncate(time.Hour)
	Printfln("Truncated  Hours: %v", trunc.Hours())
	Printfln("Rounded Mins: %v", trunc.Minutes())
}

func CreatingDurationRelativeToTime() {
	toYears := func(d time.Duration) int {
		return int(d.Hours() / (24 * 365))
	}
	future := time.Date(2051, 0, 0, 0, 0, 0, 0, time.Local)
	past := time.Date(1965, 0, 0, 0, 0, 0, 0, time.Local)
	Printfln("Future: %v", toYears(time.Until(future)))
	Printfln("Past: %v", toYears(time.Since(past)))
}

func CreatingDurationFromString() {
	d, err := time.ParseDuration("1h30m")
	if err == nil {
		Printfln("Hours: %v", d.Hours())
		Printfln("Mins: %v", d.Minutes())
		Printfln("Seconds: %v", d.Seconds())
		Printfln("Millseconds: %v", d.Milliseconds())
	} else {
		fmt.Println(err.Error())
	}
}

// IMP

// Using the time Features for Goroutine and channels

// The time package provides a small set of functions
// that are useful for working with goroutines and channels

// Although these functions are all defined in the same package, they have different uses, as demonstrated in the sections that follow.

// Putting a goroutine to sleep

func writeToChannel(channel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		time.Sleep(time.Second * 1)
	}
	close(channel)
}

func PuttingGoRoutineToSleep() {
	nameChannel := make(chan string)
	go writeToChannel(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}

// Deferring Executiion of a function

// The AfterFunc function is used to defer the execution of a function for a specified period

func writeToChannelMod(channel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		//time.Sleep(time.Second * 1)
	}
	close(channel)
}

// The first AfterFunc argument is the delay period, which is five seconds in this example. The second argument
// is the function that will be executed. In this example, I want to execute the writeToChannel function,
// but AfterFunc only accepts functions without parameters or results, and so I have to use a simple wrapper.

func DeferingExecutionOfFunction() {
	nameChannel := make(chan string)

	time.AfterFunc(time.Second*5, func() {
		writeToChannelMod(nameChannel)
	})
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}

// Receiving timed notification

func writeToChannelMod2(channel chan<- string) {
	Printfln("Waiting for initial duration...")
	<-time.After(time.Second * 2)
	Printfln("Initial duration elapsed.")
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		time.Sleep(time.Second * 1)
	}
	close(channel)
}

// The effect in this example is the same as using the Sleep function, but the difference is
// that the After function returns a channel that doesn’t block until a value is read, which means
// that a direction can be specified, additional work can be performed, and then a channel read can be performed,
// with the result that the channel will block only for the remaining part of the duration.

func ReceivingTimedNotification() {
	nameChannel := make(chan string)
	go writeToChannelMod2(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}

// Using Notifications as Timeouts in select statements

// The After function can be used with select statements to provide a timeout

func writeToChannelMod3(channel chan<- string) {
	Printfln("Waiting for initial duration...")
	<-time.After(time.Second * 2)
	Printfln("Initial duration elapsed.")
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		time.Sleep(time.Second * 3)
	}
	close(channel)
}

// The select statement will block until one of the channels is ready or until the
// timer expires. This works because the select statement will block until one of its
// channels is ready and because the After function creates a channel that blocks for a specified period.

func UsingNotificationAsTimeoutsInSelectStatements() {
	nameChannel := make(chan string)
	go writeToChannelMod3(nameChannel)
	channelOpen := true
	for channelOpen {
		Printfln("Starting channel read")
		select {
		case name, ok := <-nameChannel:
			if !ok {
				channelOpen = false
			} else {
				Printfln("Read name: %v", name)
			}
		case <-time.After(time.Second * 2):
			Printfln("Timeout")
		}
	}
}

// Stopping and Resetting Timers

// The After function is useful when you are sure that you will always
// need the timed notification. If you need the option to cancel the notification

func writeToChannelMod4(channel chan<- string) {
	timer := time.NewTimer(time.Minute * 10)
	go func() {
		time.Sleep(time.Second * 2)
		Printfln("Resetting timer")
		timer.Reset(time.Second)
	}()
	Printfln("Waiting for initial duration...")
	<-timer.C
	Printfln("Initial duration elapsed.")
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		//time.Sleep(time.Second * 3)
	}
	close(channel)
}

func StoppingAndResettingTimer() {
	nameChannel := make(chan string)
	go writeToChannelMod4(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}

// Receiving Recurring Notifications

func writeToChannelMod5(nameChannel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	tickChannel := time.Tick(time.Second)
	index := 0
	for {
		<-tickChannel
		nameChannel <- names[index]
		index++
		if index == len(names) {
			index = 0
		}
	}
}

func ReceivingRecurringNotification() {
	nameChannel := make(chan string)
	go writeToChannelMod5(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}

// The Tick function returns a channel over which Time values are sent at a specified interval

// the utility of the channel created by the Tick function isn’t the Time values sent over it, but
// the periodicity at which they are sent. In this example, the Tick function is used to create a
// channel over which values will be sent every second. The channel blocks when there is no value to
// read, which allows channels created with the Tick function to control the rate at which the
// writeToChannel function generates values.

// The Tick function is useful when an indefinite sequence of signals is required. If a fixed series of values is required

func writeToChannelMod6(nameChannel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	ticker := time.NewTicker(time.Second / 10)
	index := 0
	for {
		<-ticker.C
		nameChannel <- names[index]
		index++
		if index == len(names) {
			ticker.Stop()
			close(nameChannel)
			break
		}
	}
}

// The NewTicker function to create a Ticker that is stopped once it is no longer required.

func UsingNewTickerFunction() {
	nameChannel := make(chan string)
	go writeToChannelMod6(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}

func main() {
	// Printfln("Hello, Dates and Times")
	// CreatingTimeValues()
	// CreateFormattedTimeValues()
	// CreatingTimeWithDefinedFormat()
	// ParsingTimeValueWithString()
	// SpecifyingParsingLocation()
	// UsingLocalLocation()
	// SpecifyingTimeZoneDirectly()
	// DurationMethods()
	// CreatingDurationRelativeToTime()
	// CreatingDurationFromString()
	// PuttingGoRoutineToSleep()
	// DeferingExecutionOfFunction()
	// ReceivingTimedNotification()
	// UsingNotificationAsTimeoutsInSelectStatements()
	// StoppingAndResettingTimer()
	// ReceivingRecurringNotification()
	UsingNewTickerFunction()
}
