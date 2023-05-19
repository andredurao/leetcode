// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// nums := []int{5, 7, 7, 8, 8, 10}
	nums := []int{}
	target := 0
	result := searchRange(nums, target)
	fmt.Printf("%v\n", result)
}

func searchRange(nums []int, target int) []int {
	result := make([]int, 2)

	i := sort.SearchInts(nums, target)
	found := i < len(nums) && nums[i] == target

	if found {
		// fmt.Println("found", i)
		result[0] = searchLeftMost(nums, target)
		result[1] = searchRightMost(nums, target)
	} else {
		// fmt.Println("not found")
		return []int{-1, -1}
	}

	return result
}

func searchLeftMost(nums []int, target int) int {
	// binary search leftmost
	l := 0
	r := len(nums)
	for l < r {
		m := int(math.Floor(float64(l+r) / 2.0))
		if nums[m] < target {
			l = m + 1

		} else {
			r = m
		}
	}
	return l
}

func searchRightMost(nums []int, target int) int {
	// binary search rightmost
	l := 0
	r := len(nums)
	for l < r {
		m := int(math.Floor(float64(l+r) / 2.0))
		if nums[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	return r - 1
}
