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
		wg.Done()
	}
}

// Thread 1
func main() {
	// Waitgroup here works like a credit. When all credits are used ( .done() ) the main is concluded.
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)
	// Thread 2
	go task("A", &waitGroup)
	// Thread 3
	go task("B", &waitGroup)
	// Thread 4
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous func")
			time.Sleep(500 * time.Millisecond)
			waitGroup.Done()
		}
	}()
	waitGroup.Wait()
}
