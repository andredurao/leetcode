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

func fourSum(nums []int, target int) [][]int {
	resultMap := make(map[string][]int)
	result := make([][]int, 0)
	pairs := mapPairs(&nums)

	for sum, list1 := range pairs {
		for _, idx1 := range pairs[target-sum] {
			for _, idx2 := range list1 {
				if idx1[0] != idx2[0] && idx1[1] != idx2[1] && idx1[0] != idx2[1] && idx1[1] != idx2[0] {
					// fmt.Println(sum, -sum, idx2, idx1)
					item := []int{nums[idx2[0]], nums[idx2[1]], nums[idx1[0]], nums[idx1[1]]}

					sort.IntSlice.Sort(item)
					fmt.Println(idx1, idx2, item)
					strItem := fmt.Sprintf("%d-%d-%d-%d", item[0], item[1], item[2], item[3])
					resultMap[strItem] = item
				}
			}
		}
	}

	for _, array := range resultMap {
		result = append(result, array)
	}

	return result
}

func mapPairs(nums *[]int) map[int][][]int {
	pairs := make(map[int][][]int, 0)

	for i := 0; i < len(*nums)-1; i++ {
		for j := i + 1; j < len(*nums); j++ {
			total := (*nums)[i] + (*nums)[j]
			pairs[total] = append(pairs[total], []int{i, j})
		}
	}

	return pairs
}
