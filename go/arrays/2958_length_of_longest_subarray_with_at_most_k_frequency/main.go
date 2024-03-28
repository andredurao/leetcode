// https://leetcode.com/problems/subarray-product-less-than-k/description/

package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 1, 2, 3, 1, 2}
	nums := []int{1}

	fmt.Println(maxSubarrayLength(nums, 1))
}

func maxSubarrayLength(nums []int, k int) int {
	length := 0
	l := 0
	r := 0
	num := 0
	frequencyMap := map[int]int{}

	for r < len(nums) {
		num = nums[r]
		_, found := frequencyMap[num]
		if !found {
			frequencyMap[num] = 0
		}
		frequencyMap[num]++

		if frequencyMap[num] > k {
			length = max(length, r-l)
			// slide left until the current value has been found
			for frequencyMap[num] > k {
				frequencyMap[nums[l]]--
				l++
			}
		}
		r++
	}

	return max(length, r-l)
}
