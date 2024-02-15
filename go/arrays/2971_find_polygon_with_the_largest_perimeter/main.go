// https://leetcode.com/problems/find-polygon-with-the-largest-perimeter/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{5, 5, 5}
	fmt.Println(largestPerimeter(nums))
}

func largestPerimeter(nums []int) int64 {
	var res, sum int64

	res = -1
	sum = 0

	sort.IntSlice.Sort(nums)

	for i := range nums {
		num := int64(nums[i])
		if i >= 2 && num < sum {
			res = sum + num
		}
		sum += num
	}

	return res
}
