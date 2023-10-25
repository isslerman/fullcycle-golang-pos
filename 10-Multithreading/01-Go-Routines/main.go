package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(500 * time.Millisecond)
	}
}

// Thread 1
func main() {
	// Thread 2
	go task("A")
	// Thread 3
	go task("B")
	// Thread 4
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous func")
			time.Sleep(500 * time.Millisecond)
		}
	}()
	// We are leaving this sleep because we have the main code running on the Thread 1.
	// If we not put the sleep, the main will exit and finish the code.
	time.Sleep(15 * time.Second)
}
