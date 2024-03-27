// https://leetcode.com/problems/subarray-product-less-than-k/description/

package main

import "fmt"

func main() {
	nums := []int{10, 5, 2, 6}
	fmt.Println(numSubarrayProductLessThanK(nums, 100))
}

// From editorial added a couple conditions on l's loop
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	count := 0
	product := 1

	l := 0
	r := 0

	for r < len(nums) {
		product *= nums[r]

		for product >= k && l <= r && l < len(nums) {
			product /= nums[l]
			l++
		}

		count += r - l + 1
		r++
	}

	return count
}
