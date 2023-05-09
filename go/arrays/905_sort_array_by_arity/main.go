package main

import "fmt"

// https://leetcode.com/problems/sort-array-by-parity/

func sortArrayByParity(nums []int) []int {
	odds := make([]int, 0)
	evens := make([]int, 0)

	for _, val := range nums {
		if val%2 == 0 {
			evens = append(evens, val)
		} else {
			odds = append(odds, val)
		}
	}

	return append(evens, odds...)
}

func main() {
	nums := []int{3, 1, 2, 4}
	result := sortArrayByParity(nums)
	fmt.Println("%v", result)
}
