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
