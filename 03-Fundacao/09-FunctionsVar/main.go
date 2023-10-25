package main

import "fmt"

func main() {
	fmt.Println(sum(1, 1, 2, 4, 8, 16, 32, 64))
}

func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}
