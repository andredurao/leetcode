// https://leetcode.com/problems/find-all-duplicates-in-an-array/description/

package main

import "fmt"

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums))
}

func findDuplicates(nums []int) []int {
	res := []int{}
	numsMap := map[int]struct{}{}

	for _, num := range nums {
		_, found := numsMap[num]
		if found {
			res = append(res, num)
		} else {
			numsMap[num] = struct{}{}
		}
	}

	return res
}
