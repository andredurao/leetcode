// https://leetcode.com/problems/subarray-product-less-than-k/description/

package main

import "fmt"

func main() {
	nums := []int{10, 5, 2, 6}
	fmt.Println(numSubarrayProductLessThanK(nums, 100))
}

// From editorial
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	count := 0
	product := 1

	l := 0
	r := 0

	for ; r < len(nums); r++ {
		product *= nums[r]

		// product overflows
		for product >= k {
			product /= nums[l]
			l++
		}

		count += r - l + 1
	}

	return count
}
