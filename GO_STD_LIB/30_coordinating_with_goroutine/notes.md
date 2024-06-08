# Coordinating Goroutine

![Coordinating Goroutine](./assets/qusetions.png)

## WaitGroup

![Wait Group](./assets/waitgroup.png)

## Mutex

![Mutex](./assets/mutex.png)


### Avoiding The Mutex Pitfalls

* The best approach to using mutual exclusion is to be careful and be conservative. You must ensure that all code that accesses shared data does so using the same Mutex, and every call to a Lock method must be balanced by a call to the Unlock method. It can be tempting to try to create clever enhancements or optimizations, but doing so can lead to poor performance or application deadlocks.

## Read Write Mutex (RW Mutex)

![Read Write Mutex](./assets/RWmutex.png)

## Using Condition to coordinate goroutine

![Coordinate Goroutine](./assets/coordinate_goroutine.png)

### Cond Struct

![Cond Struct](./assets/CondStruct.png)

### Ensuring a function is executed once

![Ensuring a function is Executed once](./assets/sync%20once.png)

## Using Context

![Using Context](./assets/usingContext.png)

### Function of context package

![Function of context package](./assets/function_of_context_package.png)