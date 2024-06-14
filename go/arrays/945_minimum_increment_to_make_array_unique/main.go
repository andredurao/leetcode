// https://leetcode.com/problems/minimum-increment-to-make-array-unique/

package main

import "sort"

func main() {
	nums := []int{3, 2, 1, 2, 1, 7}
	println(minIncrementForUnique(nums))
}

func minIncrementForUnique(nums []int) int {
	sort.IntSlice.Sort(nums)

	total := 0

	for i := range nums {
		if i == 0 {
			continue
		}
		if nums[i] <= nums[i-1] {
			total += (nums[i-1] + 1) - nums[i]
			nums[i] = nums[i-1] + 1
		}
	}

	return total
}
