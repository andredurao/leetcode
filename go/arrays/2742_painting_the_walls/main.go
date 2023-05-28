// https://leetcode.com/problems/painting-the-walls/description/

package main

import (
	"fmt"
	"math"
)

func main() {
	cost := []int{1, 2, 3, 2}
	time := []int{1, 2, 3, 2}
	// cost := []int{26, 53, 10, 24, 25, 20, 63, 51}
	// time := []int{1, 1, 1, 1, 2, 2, 2, 1}
	result := paintWalls(cost, time)
	fmt.Println(result)

	// cost = []int{8, 7, 5, 15}
	// time = []int{1, 1, 2, 1}
	// result = paintWalls(cost, time)
	// fmt.Println(result)

	// cost = []int{42, 8, 28, 35, 21, 13, 21, 35}
	// time = []int{2, 1, 1, 1, 2, 1, 1, 2}
	// result = paintWalls(cost, time)
	// fmt.Println(result)
}

func paintWalls(cost []int, time []int) int {
	n := len(cost)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		dp[n][i] = int(1e9)
	}

	for i := n - 1; i >= 0; i-- {
		for remain := 1; remain <= n; remain++ {
			max := int(math.Max(0.0, float64(remain-1-time[i])))
			// max := max(0, remain-1-time[i])
			paint := cost[i] + dp[i+1][max]
			dontPaint := dp[i+1][remain]
			min := int(math.Min(float64(paint), float64(dontPaint)))
			// min := min(paint, dontPaint)
			dp[i][remain] = min
		}
	}

	return dp[0][n]
}
