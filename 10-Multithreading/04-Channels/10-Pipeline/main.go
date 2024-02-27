package main

import "fmt"

// A Pipeline is a sequence of stages
// Start -> Stage 1 -> Stage 2 -> Final

func main() {
	// input
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// stage 1
	dataChannel := sliceToChannel(nums)
	// stage 2
	doubleChannel := double(dataChannel)
	// Final
	for n := range doubleChannel {
		fmt.Print(n, " ")
	}

}

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int, 2)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func double(in <-chan int) <-chan int {
	out := make(chan int, 2)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()

	return out
}
