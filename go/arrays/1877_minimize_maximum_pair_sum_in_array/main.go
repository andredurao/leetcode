// https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/description/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 5, 2, 3}
	fmt.Println(nums)
	fmt.Println(minPairSum(nums))
}

func minPairSum(nums []int) int {
	sort.IntSlice.Sort(nums)

	max := 0

	for i := 0; i < len(nums)/2; i++ {
		sum := nums[i] + nums[len(nums)-i-1]
		if sum > max {
			max = sum
		}
	}

	return max
}
