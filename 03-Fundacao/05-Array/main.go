package main

import "fmt"

// const a = "Hello Boss"

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 5

	fmt.Println(meuArray[0])
	fmt.Println(len(meuArray))

	for i, v := range meuArray {
		fmt.Printf("O valor do indice [%d] é %d\n", i, v)
	}
}
