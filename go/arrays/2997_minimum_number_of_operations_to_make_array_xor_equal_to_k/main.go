// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-xor-equal-to-k/

package main

func main() {
	nums := []int{2, 1, 3, 4}
	k := 1
	println(minOperations(nums, k))
}

func minOperations(nums []int, k int) int {
	xor := 0

	for i := 0; i < len(nums); i++ {
		xor = xor ^ nums[i]
	}

	total := 0

	for xor > 0 || k > 0 {
		if (xor & 1) != (k & 1) {
			total++
		}
		xor >>= 1
		k >>= 1
	}

	return total
}
