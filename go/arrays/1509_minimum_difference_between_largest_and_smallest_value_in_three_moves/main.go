// https://leetcode.com/problems/minimum-difference-between-largest-and-smallest-value-in-three-moves

package main

import (
	"math"
	"sort"
)

func main() {
	nums := []int{1, 5, 0, 10, 14}
	res := minDifference(nums)
	println(res)
}

func minDifference(nums []int) int {
	sort.IntSlice.Sort(nums)
	if len(nums) <= 3 {
		return 0
	}

	res := math.MaxInt

	for i := 0; i <= 3; i++ {
		res = min(res, abs(nums[i]-nums[len(nums)-4+i]))
	}

	return res
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
