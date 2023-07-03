package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{2, 2, 1, 2, 1}
	res := maximumElementAfterDecrementingAndRearranging(arr)
	fmt.Println(res)
}

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	total := 1

	sort.IntSlice.Sort(arr)
	for i := 1; i < len(arr); i++ {
		if arr[i] >= (total + 1) {
			total++
		}
	}

	return total
}
