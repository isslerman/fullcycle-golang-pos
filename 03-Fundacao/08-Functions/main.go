package main

import (
	"errors"
	"fmt"
)

var globalVar int

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum2(10, 2))
	fmt.Println(sum2(10, 200))

	soma, err := sum3(100, 200)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(soma)
	globalVar = 10
	point := &globalVar
	point2 := &point
	fmt.Printf("Globalvar: %d, add:%p, type: %T\n", globalVar, &globalVar, globalVar)
	fmt.Printf("point: %v, value: %v, add:%p, type: %T\n", point, *point, &point, point)
	fmt.Printf("point2: %v, value: %v, add:%p, type: %T\n", point2, **point2, &point2, point2)
	printPointers(point)

	changeGlobalVar()
	fmt.Printf("Globalvar: %d, add:%p, type: %T\n", globalVar, &globalVar, globalVar)
}

func printPointers(p *int) {
	fmt.Printf("PrintPoint: %v, %T\n", *p, p)
}

func changeGlobalVar() {
	globalVar = 50
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
