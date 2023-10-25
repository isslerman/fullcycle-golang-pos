package main

import (
	"fmt"

	"github.com/isslerman/goexpert/07/01/math"
)

func main() {
	fmt.Println("Oi Boss")

	m := math.Exported_Math{A: 10, B: 20}
	fmt.Println(m.Add())
	fmt.Println(math.NewAdd(10, 20))
}
