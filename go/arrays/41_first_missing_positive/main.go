// https://leetcode.com/problems/first-missing-positive/

package main

import "fmt"

func main() {
	nums := []int{1, 0, 2}
	fmt.Println(firstMissingPositive(nums))
}

func firstMissingPositive(nums []int) int {
	positivesMap := map[int]struct{}{}
	maxNum := 0

	for _, num := range nums {
		if num > 0 {
			maxNum = max(maxNum, num)
			positivesMap[num] = struct{}{}
		}
	}

	for i := 1; i < maxNum; i++ {
		_, found := positivesMap[i]
		if !found {
			return i
		}
	}

	return maxNum + 1
}
