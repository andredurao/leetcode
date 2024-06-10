// https://leetcode.com/problems/height-checker/

package main

import "sort"

func main() {
	heights := []int{1, 1, 4, 2, 1, 3}
	println(heightChecker(heights))
}

func heightChecker(heights []int) int {
	total := 0

	sortedHeights := make([]int, len(heights))
	copy(sortedHeights, heights)
	sort.IntSlice.Sort(sortedHeights)

	for i := range heights {
		if heights[i] != sortedHeights[i] {
			total++
		}
	}

	return total
}
