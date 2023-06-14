// https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/description/
package main

import (
	"fmt"
	"sort"
)

func main() {
	// arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	arr := []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
	result := sortByBits(arr)
	fmt.Printf("%v\n", result)
}

func sortByBits(arr []int) []int {
	aux := make([][]int, len(arr))
	for i, val := range arr {
		aux[i] = []int{val, bitCountK(val)}
	}

	sort.Slice(aux, func(i, j int) bool { return aux[i][1] < aux[j][1] || (aux[i][1] == aux[j][1] && aux[i][0] < aux[j][0]) })
	for i := range arr {
		arr[i] = aux[i][0]
	}

	return arr
}

// kernighan's bit count (from: https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetKernighan)
func bitCountK(n int) int {
	c := 0
	for n > 0 {
		n = n & (n - 1)
		c++
	}
	return c
}

func bitCount(n int) (result int) {
	for n > 0 {
		result = result + (n & 1)
		n = n >> 1
	}
	return
}
