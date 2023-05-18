// https://leetcode.com/problems/majority-element-ii/description/

package main

import (
	"fmt"
)

func main() {
	nums := []int{3}
	result := majorityElement(nums)
	fmt.Printf("%v\n", result)
}

func majorityElement(nums []int) []int {
	result := make([]int, 0)
	counts_map := make(map[int]int)
	threshold := len(nums) / 3
	// fmt.Println("threshold", threshold)

	for _, num := range nums {
		val, found := counts_map[num]
		if !found {
			counts_map[num] = 1
		} else {
			counts_map[num]++
		}

		val, _ = counts_map[num]

		if val > threshold {
			counts_map[num] = -1
			result = append(result, num)
		}
	}

	return result
}
