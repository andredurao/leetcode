// https://leetcode.com/problems/132-pattern/description/

package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Printf("%v\n", find132pattern([]int{1, 2, 3, 4}))

	fmt.Printf("%v\n", find132pattern([]int{3, 1, 4, 2}))
}

// based on solution
func find132pattern(nums []int) bool {
	size := len(nums)
	if size < 3 {
		return false
	}

	mins := make([]int, size)
	mins[0] = nums[0]

	for i := 1; i < size; i++ {
		mins[i] = int(math.Min(float64(mins[i-1]), float64(nums[i])))
	}

	k := size
	for j := size - 1; j >= 0; j-- {
		if nums[j] <= mins[j] {
			continue
		}
		for k < size && nums[k] <= mins[j] {
			k++
		}
		if k < size && nums[k] < nums[j] {
			return true
		}
		k--
		nums[k] = nums[j]
	}

	fmt.Printf("%v\n", mins)

	return false
}

func is_132(nums []int, i int, j int, k int) bool {
	return (i < j && j < k) && (nums[i] < nums[k] && nums[k] < nums[j])
}
