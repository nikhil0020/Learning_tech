package main

import (
	"context"
	"math"
	"math/rand"
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{}

// Run the below code to see why mutual exclusion is important
// This code will give different output because goroutine are using same counter value
// Sometime the result is 15000 which is correct but most of the time result is not 15000

/*
func doSum(count int, val *int) {
	time.Sleep(time.Second)
	for i := 0; i < count; i++ {
		*val++
	}
	waitGroup.Done()
}
func main() {
	counter := 0
	numRoutines := 3
	waitGroup.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go doSum(5000, &counter)
	}
	waitGroup.Wait()
	Printfln("Total: %v", counter)
}

*/

func doSum(count int, val *int) {
	for i := 0; i < count; i++ {
		*val++
	}
	waitGroup.Done()
}

func MutexLock() {
	counter := 0
	waitGroup.Add(1)
	go doSum(5000, &counter)
	waitGroup.Wait()
	Printfln("Total: %v", counter)
}

// Read Write Mutex

var rwmutex = sync.RWMutex{}
var squares = map[int]int{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func calculateSquares(max, iterations int) {
	defer waitGroup.Done()
	for i := 0; i < iterations; i++ {
		val := rnd.Intn(max)
		rwmutex.RLock()
		square, ok := squares[val]
		rwmutex.RUnlock()
		if ok {
			Printfln("Cached value: %v = %v", val, square)
		} else {
			rwmutex.Lock()
			if _, ok := squares[val]; !ok {
				squares[val] = int(math.Pow(float64(val), 2))
				Printfln("Added value: %v = %v", val, squares[val])
			}
			rwmutex.Unlock()
		}
	}
}

func ReadWriteMutex() {
	numRoutines := 3
	waitGroup.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go calculateSquares(10, 5)
	}
	waitGroup.Wait()
	Printfln("Cached values: %v", len(squares))
}

// Conditional Goroutines

var readyCond = sync.NewCond(rwmutex.RLocker())

func generateSquares(max int) {
	rwmutex.Lock()
	Printfln("Generating data...")
	for val := 0; val < max; val++ {
		squares[val] = int(math.Pow(float64(val), 2))
	}
	rwmutex.Unlock()
	Printfln("Broadcasting condition")
	readyCond.Broadcast()
	waitGroup.Done()
}
func readSquares(id, max, iterations int) {
	readyCond.L.Lock()
	for len(squares) == 0 {
		readyCond.Wait()
	}
	for i := 0; i < iterations; i++ {
		key := rnd.Intn(max)
		Printfln("#%v Read value: %v = %v", id, key, squares[key])
		time.Sleep(time.Millisecond * 10)
	}
	readyCond.L.Unlock()
	waitGroup.Done()
}

func SyncCond() {
	numRoutines := 2
	waitGroup.Add(numRoutines)

	for i := 0; i < numRoutines; i++ {
		go readSquares(i, 10, 5)
	}

	waitGroup.Add(1)

	go generateSquares(10)

	waitGroup.Wait()

	Printfln("Cached values: %v", len(squares))
}

// Ensuring a function is executed once

var once = sync.Once{}

func generateSquaresOnce(max int) {
	//rwmutex.Lock()
	Printfln("Generating data...")
	for val := 0; val < max; val++ {
		squares[val] = int(math.Pow(float64(val), 2))
	}
	// rwmutex.Unlock()
	// Printfln("Broadcasting condition")
	// readyCond.Broadcast()
	// waitGroup.Done()
}
func readSquaresOnce(id, max, iterations int) {
	once.Do(func() {
		generateSquaresOnce(max)
	})
	// readyCond.L.Lock()
	// for len(squares) == 0 {
	//     readyCond.Wait()
	// }
	for i := 0; i < iterations; i++ {
		key := rand.Intn(max)
		Printfln("#%v Read value: %v = %v", id, key, squares[key])
		time.Sleep(time.Millisecond * 100)
	}
	//readyCond.L.Unlock()
	waitGroup.Done()
}

func SyncOnce() {
	rnd.Seed(time.Now().UnixNano())
	numRoutines := 2
	waitGroup.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go readSquaresOnce(i, 10, 5)
	}
	// waitGroup.Add(1)
	// go generateSquares(10)
	waitGroup.Wait()
	Printfln("Cached values: %v", len(squares))
}

// Using the Once struct simplifies the example because the Do method blocks until
// the function it receives has been executed, after which it returns without executing
// the function again. Since the only changes to the shared data in this example are made by
// the generateSquares function, using the Do method to execute this function ensures that the
// changes are made safely. Not all code fits the Once model so well, but in this example, I can
// remove the RWMutex and the Cond. Compile and execute the project, and you will see output similar to the following:

func processRequest(ctx context.Context, wg *sync.WaitGroup, count int) {
	total := 0
	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			Printfln("Stopping processing - request cancelled")
			goto end
		default:
			Printfln("Processing request: %v", total)
			total++
			time.Sleep(time.Millisecond * 250)
		}
	}

	Printfln("Request processed....%v", total)
end:
	wg.Done()
}

func CancellingRequest() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	Printfln("Request dispatched...")
	ctx, cancel := context.WithCancel(context.Background())
	go processRequest(ctx, &wg, 10)
	time.Sleep(2 * time.Second)
	Printfln("Cancelling Request")
	cancel()
	wg.Wait()
}

// Setting A deadline

func processRequestDeadline(ctx context.Context, wg *sync.WaitGroup, count int) {
	total := 0
	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.Canceled {
				Printfln("Stopping processing - request cancelled")
			} else {
				Printfln("Stopping processing - deadline reached")
			}
			goto end
		default:
			Printfln("Processing request: %v", total)
			total++
			time.Sleep(time.Millisecond * 250)
		}
	}
	Printfln("Request processed... %v", total)
end:
	wg.Done()
}

func DeadlineReachedRequest() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	Printfln("Request dispatched...")
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	go processRequestDeadline(ctx, &wg, 10)
	// time.Sleep(time.Second)
	// cancel()
	wg.Wait()
}

// Providing Request Data

const (
	countKey = iota
	sleepPeriodKey
)

func processRequestWithData(ctx context.Context, wg *sync.WaitGroup) {
	total := 0
	count := ctx.Value(countKey).(int)
	sleepPeriod := ctx.Value(sleepPeriodKey).(time.Duration)

	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.Canceled {
				Printfln("Stopping processing - request cancelled")
			} else {
				Printfln("Stopping processing - deadline cancelled")
			}
			goto end
		default:
			Printfln("Request processed...%v", total)
			total++
			time.Sleep(sleepPeriod)
		}
	}
	Printfln("Request processed...%v", total)
end:
	wg.Done()
}

func ProvidingRequestData() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	Printfln("Request dispatched...")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	ctx = context.WithValue(ctx, countKey, 10)
	ctx = context.WithValue(ctx, sleepPeriodKey, time.Millisecond*399)
	go processRequestWithData(ctx, &wg)
	wg.Wait()
}

func main() {
	// MutexLock()
	// ReadWriteMutex()
	// SyncCond()
	// SyncOnce()
	// CancellingRequest()
	// DeadlineReachedRequest()
	ProvidingRequestData()
}
