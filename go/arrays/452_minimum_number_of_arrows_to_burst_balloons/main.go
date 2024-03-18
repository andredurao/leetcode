// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	fmt.Println(findMinArrowShots(points))
}

// solution from Editorial

func findMinArrowShots(points [][]int) int {
	slices.SortFunc(points, func(a, b []int) int {
		return cmp.Compare(a[1], b[1])
	})

	stackEnd := points[0][1]
	total := 1

	for _, point := range points {
		if point[0] > stackEnd {
			total++
			stackEnd = point[1]
		}
	}

	return total
}
