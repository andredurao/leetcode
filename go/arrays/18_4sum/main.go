// https://leetcode.com/problems/4sum/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	result := fourSum(nums, target)
	fmt.Printf("%v\n", result)
}

// From editorial
func fourSum(nums []int, target int) [][]int {
	sort.IntSlice.Sort(nums)
	return kSum(nums, target, 4)
}

func kSum(nums []int, target int, k int) [][]int {
	result := make([][]int, 0)

	if len(nums) == 0 {
		return result
	}

	average_value := target / k
	if average_value < nums[0] || nums[len(nums)-1] < average_value {
		return result
	}

	if k == 2 {
		return twoSum(nums, target)
	}

	for i := range nums {
		if i == 0 || nums[i-1] != nums[i] {
			kSumResult := kSum(nums[i+1:], target-nums[i], k-1)
			for _, subsets := range kSumResult {
				subsets = append([]int{nums[i]}, subsets...)
				result = append(result, subsets)
			}
		}
	}

	return result
}

func twoSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	lo := 0
	hi := len(nums) - 1

	for lo < hi {

		sum := nums[lo] + nums[hi]
		if sum < target || (lo > 0 && nums[lo] == nums[lo-1]) {
			lo++
		} else if sum > target || (hi < len(nums)-1 && nums[hi] == nums[hi+1]) {
			hi--
		} else {
			result = append(result, []int{nums[lo], nums[hi]})
			lo++
			hi--
		}
	}
	return result
}
