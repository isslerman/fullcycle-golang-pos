package main

import "fmt"

func main() {

	m1 := [][]int{
		{1, 1},
		{1, 1},
	}

	m2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	sl, sr := sumDiagonals(m1)
	fmt.Printf("The Left and the Right sum is: %d, %d\n", sl, sr)
	sl, sr = sumDiagonals(m2)
	fmt.Printf("The Left and the Right sum is: %d, %d\n", sl, sr)
}

// sumDiagonals calculates the sum of both the left and right diagonal elements of a square matrix.
func sumDiagonals(matrix [][]int) (int, int) {
	sumLeft := 0
	sumRight := 0
	n := len(matrix)
	for i := 0; i < n; i++ {
		// Left diagonal: row index equals column index
		sumLeft += matrix[i][i]
		// Right diagonal: row index plus column index equals matrix size minus one
		sumRight += matrix[i][n-i-1]
	}
	return sumLeft, sumRight
}
