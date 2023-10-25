package main

import "fmt"

func main() {

	total := func() int {
		return sum(1, 1, 2, 4, 8, 16, 32, 64) * 2
	}()

	total2 := sum(1, 1, 2, 4, 8, 16, 32, 64) * 2

	fmt.Println(total)
	fmt.Println(total2)
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}
