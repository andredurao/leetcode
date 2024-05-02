// https://leetcode.com/problems/largest-positive-integer-that-exists-with-its-negative/

package main

func main() {
	nums := []int{-1, 2, -3, 3}
	println(findMaxK(nums))
}

func findMaxK(nums []int) int {
	numsMap := map[int]struct{}{}

	for _, num := range nums {
		numsMap[num] = struct{}{}
	}

	maxK := -1
	for _, num := range nums {
		_, found := numsMap[-num]
		if num > maxK && found {
			maxK = num
		}
	}

	return maxK
}
