package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// taskA performs some work and signals completion
func taskA() {
	defer wg.Done() // Decrement the waitgroup counter when done
	fmt.Println("Starting Task A...")
	time.Sleep(2 * time.Second) // Simulate work
	fmt.Println("Task A completed!")
}

// taskB waits for taskA to finish and then starts its work
func taskB() {
	defer wg.Done() // Decrement the waitgroup counter when done
	fmt.Println("Waiting for Task A to complete...")

	// Simulate some dependency on Task A's completion
	time.Sleep(1 * time.Second)
	fmt.Println("Task B can start now after Task A.")
	time.Sleep(1 * time.Second) // Simulate work
	fmt.Println("Task B completed!")
}

func main() {
	wg.Add(1) // We have two tasks to wait for

	go taskA() // Launch Task A in a goroutine
	wg.Add(1)  // We have two tasks to wait for

	go taskB() // Launch Task B in a goroutine
	wg.Wait()

	// Wait for all tasks to complete
	fmt.Println("All tasks completed!")

}
