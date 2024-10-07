package main

import "fmt"

func main() {
	c := make(chan int)
	go func(i int) {
		result := i * 5
		c <- result
	}(5)
	result := <-c

	my := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	removeindex := 5

	my = append(my[:removeindex], my[removeindex+1:]...)

	fmt.Println(result)
	fmt.Println(my)
	println(len(my))

	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // Close the channel
	}()

	for value := range ch { // Receive values until the channel is closed
		fmt.Println(value)
	}
	fmt.Scanln()
}
