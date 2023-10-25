package main

import "fmt"

// Thread 1
func main() {
	channel := make(chan string) // empty

	// Thread 2
	go func() {
		channel <- "Yes Boss!" // full
	}()

	msg := <-channel // Channel empty again
	fmt.Println(msg)
}
