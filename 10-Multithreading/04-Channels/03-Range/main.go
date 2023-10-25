package main

import "fmt"

// Thread 1
func main() {
	ch := make(chan int)
	go publish(ch)
	reader(ch)
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// we need to put the close here to not get an deadlock.
	close(ch)
}

// it keeps reading until the channel close
func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}
