// https://leetcode.com/problems/binary-subarrays-with-sum/description/

package main

import "fmt"

func main() {
	nums := []int{0, 0, 0, 0, 0, 0, 1, 0, 0, 0}
	goal := 0
	fmt.Println(numSubarraysWithSum(nums, goal))
}

func numSubarraysWithSum(nums []int, goal int) int {
	if len(nums) < goal || len(nums) == 0 {
		return 0
	}

	res := 0
	total := 0

	for i := 0; i < len(nums); i++ {
		total = nums[i]
		if total == goal {
			res++
		}
		if total <= goal {
			for j := i + 1; j < len(nums); j++ {
				total += nums[j]
				if total == goal {
					res++
				} else if total > goal {
					break
				}
			}
		}
	}

	return res
}
