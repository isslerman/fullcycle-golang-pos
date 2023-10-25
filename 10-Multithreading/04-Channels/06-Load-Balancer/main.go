package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int, speed int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}

func main() {
	data := make(chan int)
	QtdWorkers := 10000

	for i := 0; i < QtdWorkers; i++ {
		go worker(i, data, 5)
	}

	for i := 0; i < 1000000; i++ {
		data <- i
	}
}
