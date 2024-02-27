package main

import (
	"fmt"
)

func main() {
	a := []string{"", "a", "bf", "ba", "aidud"}
	b := []string{"", "fc", "ba", "bbb", "dsikd"}

	totalChecks := len(a)
	result := make([]int32, 0, len(a))

	// check over size of array
	for i := 0; i < totalChecks; i++ {
		if len(a) == 0 || len(b) == 0 {
			result = append(result, -1)
			break
		}
		word1 := a[i]
		word2 := b[i]
		// fmt.Printf("First letter of word1 is %v", string(word1[0]))
		// fmt.Printf("First letter of word2 is %v", string(word2[0]))

		sA := len(word1)
		sB := len(word2)
		// fmt.Printf("Sa: %d Sb: %d\n", sA, sB)

		// check first word of each array
		// if size is diff we can return -1
		if sA != sB {
			fmt.Println("Result: -1")
			result = append(result, -1)
			// if size is equal
		} else {
			counter := 0
			// equal size, pal possible
			fmt.Printf("W1: %v W2: %v\n", word1, word2)
			for i, j := 0, sB-1; i < sA; i, j = i+1, j-1 {
				fmt.Printf("Letters: %v-%v\n", string(word1[i]), string(word2[j]))
				if word1[i] != word2[j] {
					counter++
					fmt.Println(counter)
				}
			}
			result = append(result, int32(counter))
			fmt.Println("Total de mudancas", counter)
		}
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("Final Result:", result)
	}
}
