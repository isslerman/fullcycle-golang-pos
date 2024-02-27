package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(500 * time.Millisecond)
	}
	wg.Done()
}

// Thread 1
func main() {
	// Waitgroup here works like a credit. When all credits are used ( .done() ) the main is concluded.
	waitGroup := sync.WaitGroup{}
	// Thread 2
	waitGroup.Add(1)
	go task("A", &waitGroup)
	// Thread 3
	waitGroup.Add(1)
	go task("B", &waitGroup)
	// Thread 4
	go func() {
		waitGroup.Add(1)
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous func")
			time.Sleep(500 * time.Millisecond)
		}
		waitGroup.Done()
	}()
	waitGroup.Wait()
}
