package main

import (
	"fmt"
	"time"
)

func main() {
	chars := []string{"a", "bb", "bc", "dd"}
	ch := make(chan string, 4)

	for _, s := range chars {
		// select {
		// case ch <- s:
		// }
		ch <- s
	}
	close(ch)

	for result := range ch {
		fmt.Printf("%s ", result)
	}

	example2()
}

func example2() {
	done := make(chan bool)
	go doWork(done)

	time.Sleep(time.Second * 3)
	close(done)

}

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Working hard here. Make it stop, please.")
		}
	}
}
