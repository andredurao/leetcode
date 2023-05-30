// https://leetcode.com/problems/pascals-triangle-ii/

package main

import "fmt"

func main() {
	result := getRow(3)
	fmt.Printf("%v\n", result)
}

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	result := make([]int, rowIndex+1)
	// https://en.wikipedia.org/wiki/Pascal%27s_triangle#Formula
	n := rowIndex
	for k := 0; k <= n; k++ {
		if k == 0 || k == n {
			result[k] = 1
		} else {
			result[k] = result[k-1] * (n + 1 - k) / k
		}

	}
	return result
}
