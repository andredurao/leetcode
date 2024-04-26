// https://leetcode.com/problems/minimum-falling-path-sum-ii/description/

package main

import "math"

func main() {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	println(minFallingPathSum(grid))
}

func minFallingPathSum(grid [][]int) int {
	dp := [][]int{}
	for row := range grid {
		if row == 0 {
			dp = append(dp, grid[0])
		} else {
			dp = append(dp, make([]int, len(grid[row])))
		}
	}

	// second row
	minPrev, minPrevIndex := minValue(&dp[0], -1)
	if len(grid) == 1 {
		return minPrev
	}

	for i := 1; i < len(grid); i++ {
		for j := range grid[1] {
			if minPrevIndex == j {
				cMinValue, _ := minValue(&dp[i-1], j)
				dp[i][j] = grid[i][j] + cMinValue
			} else {
				dp[i][j] = grid[i][j] + minPrev
			}
		}
		minPrev, minPrevIndex = minValue(&dp[i], -1)
	}

	res, _ := minValue(&dp[len(dp)-1], -1)

	return res
}

func minValue(row *[]int, ignore int) (int, int) {
	res := math.MaxInt32
	index := -1
	for i, val := range *row {
		if i == ignore {
			continue
		}
		if val < res {
			res = val
			index = i
		}
	}

	return res, index
}
