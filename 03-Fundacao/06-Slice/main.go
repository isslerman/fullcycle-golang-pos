package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 4, 8, 16, 32, 64, 128}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// pegando primeiros 0 itens
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	// pegando primeiros 4 itens
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])
	// cortando os dois primeiros e ficando com o restante. descartando os primeiros
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 256) // aqui a capacidade do array é dobrada e não incrementada em 1.
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	min, max := MinMax(s)
	fmt.Printf("Min: %d, Max: %d\n", min, max)
	fmt.Printf("Min: %d\n", Min(s))
	fmt.Printf("Max: %d\n", Max(s))
	fmt.Printf("MinNil: %d\n", Min([]int{}))

}

func MinMax(array []int) (int, int) {
	if len(array) == 0 {
		panic("empty array")
	}
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func Min(array []int) int {
	if len(array) == 0 {
		panic("empty array")
	}

	var min int = array[0]
	for _, value := range array {
		if min > value {
			min = value
		}
	}
	return min
}

func Max(array []int) int {
	if len(array) == 0 {
		panic("empty array")
	}
	var max int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
	}
	return max
}

func removeValue(A []int32, v int32) []int32 {
	var result []int32
	for i, a := range A {
		if v == a {
			return append(A[:i], A[i+1:]...)
		}
	}
	return result
}

func rotateLeft(d int32, arr []int32) []int32 {
	if d == 0 || len(arr) == 0 {
		return arr
	}
	shifted := make([]int32, len(arr))

	for i := range arr {
		shifted[i] = arr[(int32(int32(i)+d))%int32(len(arr))]
	}
	return shifted
}
