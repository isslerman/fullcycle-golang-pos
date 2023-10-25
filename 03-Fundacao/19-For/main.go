package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "tres"}
	for k, v := range numeros {
		println(k, v)
	}

	for _, v := range numeros {
		println(v)
	}

	for k, _ := range numeros {
		println(k)
	}

	j := 0
	for j < 10 {
		println(j)
		j++
	}

	count := 0
	for {
		fmt.Printf("sou eterno %v\n", count)
		count++
	}
}
