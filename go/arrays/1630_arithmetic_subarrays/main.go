// https://leetcode.com/problems/arithmetic-subarrays/description/

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4, 6, 5, 9, 3, 7}
	l := []int{0, 0, 2}
	r := []int{2, 3, 5}
	res := checkArithmeticSubarrays(nums, l, r)
	fmt.Printf("%v\n", res)
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	result := make([]bool, len(l))

	for i := 0; i < len(l); i++ {
		subarray := make([]int, (r[i] - l[i] + 1))
		if r[i] < len(nums)-1 {
			copy(subarray, nums[l[i]:(r[i]+1)])
		} else {
			copy(subarray, nums[l[i]:])
		}
		sort.IntSlice.Sort(subarray)
		// fmt.Println(i, subarray)
		result[i] = isArithmeticArray(&subarray)
	}

	return result
}

func isArithmeticArray(subarray *[]int) bool {
	diff := 0
	for i := 1; i < len(*subarray); i++ {
		if i == 1 {
			diff = (*subarray)[i] - (*subarray)[i-1]
			continue
		}
		if ((*subarray)[i] - (*subarray)[i-1]) != diff {
			return false
		}
	}
	return true
}
