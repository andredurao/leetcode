package main

import (
	"fmt"
	"slices"
)

func main() {
	input := []int{-1, 0, 1, 2, -1, -4} // output [[-1,-1,2],[-1,0,1]]
	result := threeSum(input)
	fmt.Println(result)
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			result = append(result, twoSum(nums, i)...)
		}
	}
	return result
}

func twoSum(nums []int, i int) [][]int {
	result := make([][]int, 0)

	seen := make(map[int]struct{})
	j := i + 1
	for j < len(nums) {
		complement := -nums[i] - nums[j]
		_, found := seen[complement]
		if found {
			result = append(result, []int{nums[i], nums[j], complement})
			for j+1 < len(nums) && nums[j] == nums[j+1] {
				j++
			}
		}
		seen[nums[j]] = struct{}{}
		j++
	}
	return result
}
