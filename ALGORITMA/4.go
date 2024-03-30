package main

import (
	"fmt"
)

func diagonalSumDifference(matrix [][]int) int {
	primarySum := 0
	secondarySum := 0
	n := len(matrix)

	for i := 0; i < n; i++ {
		primarySum += matrix[i][i] // Add primary diagonal elements
		secondarySum += matrix[i][n-i-1] // Add secondary diagonal elements
	}

	return primarySum - secondarySum
}

func main() {
	matrix := [][]int{{1, 2, 0}, {4, 5, 6}, {7, 8, 9}}
	difference := diagonalSumDifference(matrix)
	fmt.Println("Hasil pengurangan diagonal adalah:", difference) // Output: 3
}
