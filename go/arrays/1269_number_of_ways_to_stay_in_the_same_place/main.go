// https://leetcode.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/

package main

import (
	"fmt"
	"math"
)

func main() {
	steps := 3
	arrLen := 2
	result := numWays(steps, arrLen)
	fmt.Println(result)
}

func numWays(steps int, arrLen int) int {
	minArrLen := int(math.Min(float64(steps), float64(arrLen)))
	memo := make([][]int, minArrLen)
	for i := 0; i < minArrLen; i++ {
		memo[i] = make([]int, steps+1)
		for j := 0; j <= steps; j++ {
			memo[i][j] = -1
		}
	}

	result := dp(0, steps, minArrLen, &memo)
	return result
}

// Based on leetcode solution
func dp(pos int, remain int, minArrLen int, memo *[][]int) int {
	if remain == 0 {
		if pos == 0 {
			return 1
		}
		return 0
	}

	// if that value has been calculated before, just return it's result
	if (*memo)[pos][remain] != -1 {
		return (*memo)[pos][remain]
	}

	result := dp(pos, remain-1, minArrLen, memo)

	mod := int(1e9) + 7

	// left
	if pos > 0 {
		result = (result + dp(pos-1, remain-1, minArrLen, memo)) % mod
	}

	// right
	if pos < minArrLen-1 {
		result = (result + dp(pos+1, remain-1, minArrLen, memo)) % mod
	}

	(*memo)[pos][remain] = result
	return result
}
