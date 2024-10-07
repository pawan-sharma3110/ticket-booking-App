package main

import "sync"

var wg sync.WaitGroup

func A(count int) {
	for i := 0; i < count; i++ {
		println(i, ": A")
	}
	wg.Done()
}

func B(count int) {
	for i := 0; i < count; i++ {
		println(i, ": B")
	}
	wg.Done()

}
func main() {
	wg.Add(2)
	go A(50)
	go B(10)
	wg.Wait()

}
