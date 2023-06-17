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
	sumMap := make(map[int](map[string][]int))
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			sum := nums[i] + nums[j]
			_, found := sumMap[sum]
			if !found {
				sumMap[sum] = make(map[string][]int)
			}
			pair := []int{nums[i], nums[j]}
			slices.Sort(pair)
			strPair := fmt.Sprintf("%d,%d", pair[0], pair[1])
			pair = append(pair, []int{i, j}...)
			sumMap[sum][strPair] = pair
		}
	}

	resultMap := make(map[string][]int, 0)

	for i, n := range nums {
		pairs, found := sumMap[-n]
		if found {
			for _, values := range pairs {
				if i != values[2] && i != values[3] {
					item := []int{n, values[0], values[1]}
					slices.Sort(item)
					strItem := fmt.Sprintf("%d,%d,%d", item[0], item[1], item[2])
					resultMap[strItem] = item
				}
			}
		}
	}

	result := make([][]int, 0)
	for _, item := range resultMap {
		result = append(result, item)
	}
	return result
}
