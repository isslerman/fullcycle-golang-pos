package main

import "fmt"

// channel only receives data
// insert in channel
// receives only
func recebe(nome string, hello chan<- string) {
	hello <- nome
}

// read and empty the channel
// remove from channel
// send-only channel
func ler(data <-chan string) {
	fmt.Println(<-data)
}

// Thread 01
func main() {
	hello := make(chan string)
	go recebe("Hello", hello)
	ler(hello)
}
