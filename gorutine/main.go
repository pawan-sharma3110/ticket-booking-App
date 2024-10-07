package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello from Goroutine!")

	}
	wg.Done()
}

func main() {
	// Launching a goroutine
	wg.Add(1)
	go sayHello()

	fmt.Println("Main Goroutine Finished")
	wg.Wait()
}
