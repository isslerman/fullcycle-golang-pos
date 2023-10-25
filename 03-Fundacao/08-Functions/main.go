package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum2(10, 2))
	fmt.Println(sum2(10, 200))

	soma, err := sum3(100, 200)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(soma)
}

// func sum(a int, b int) int {
func sum(a, b int) int {
	return a + b
}

// two returns
func sum2(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}
	return a + b, false
}

// using error
func sum3(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("err: valor muito alto")
	}
	return a + b, nil
}
