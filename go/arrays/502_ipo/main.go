// https://leetcode.com/problems/ipo/

package main

import (
	"fmt"
)

func main() {
	k := 2
	w := 0
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 1}
	res := findMaximizedCapital(k, w, profits, capital)
	fmt.Println(res)
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	dp := [][]int{}

	for i := 0; i <= k+1; i++ {
		dp = append(dp, make([]int, k+1))
	}

	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	// this looks like a knapsack problem of sorts, because they are looking to maximize the profit
	// but with a difference: the capacity(k) is fixed, but items can only be used if there is enough capital
	// https://en.wikipedia.org/wiki/Knapsack_problem
	total := knapsack(k, w, &profits, &capital, len(profits)-1, &dp)

	return total
}

func knapsack(k int, w int, profits *[]int, capital *[]int, index int, dp *[][]int) int {
	if index < 0 {
		return 0
	}
	if (*dp)[index][k] != -1 || k == 0 {
		return (*dp)[index][k]
	}
	// doesnt fit
	if w > (*capital)[index] {
		(*dp)[index][k] = knapsack(k, w, profits, capital, index-1, dp)
	} else {
		(*dp)[index][k] = max(
			(*profits)[index]+knapsack(k-1, w, profits, capital, index-1, dp),
			knapsack(k, w, profits, capital, index-1, dp),
		)
	}
	return (*dp)[index][k]
}
