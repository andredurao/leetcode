package main

import "fmt"

// https://leetcode.com/problems/monotonic-array/

func main() {
	nums := []int{1, 2, 2, 3}
	result := isMonotonic(nums)
	fmt.Printf("%v\n", result)
}

func isMonotonic(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	// 0: neutral (initial), 1: increasing, -1: decreasing
	last_compare := 0

	for i := 1; i < len(nums); i++ {
		fmt.Printf("%v\n", i)
		diff := compare(nums[i-1], nums[i])
		if diff != 0 {
			if last_compare == 0 {
				last_compare = diff
			} else if last_compare != diff {
				return false
			}
		}
	}

	return true
}

func compare(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}
