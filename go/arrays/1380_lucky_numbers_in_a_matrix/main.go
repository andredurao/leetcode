//

package main

import (
	"fmt"
	"math"
)

func main() {
	// matrix := [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}
	matrix := [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}
	fmt.Println(luckyNumbers(matrix))
}

func luckyNumbers(matrix [][]int) []int {
	mins := make([]int, len(matrix))
	for i := range mins {
		mins[i] = math.MaxInt32
	}
	maxs := make([]int, len(matrix[0]))
	for i := range maxs {
		maxs[i] = 0
	}

	for i, row := range matrix {
		for j, n := range row {
			mins[i] = min(mins[i], n)
			maxs[j] = max(maxs[j], n)
		}
	}

	res := []int{}
	for i, row := range matrix {
		for j, n := range row {
			if mins[i] == n && maxs[j] == n {
				res = append(res, n)
			}
		}
	}

	return res
}
