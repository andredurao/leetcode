// https://leetcode.com/problems/climbing-stairs/?envType=daily-question&envId=2024-01-18

// the easiest solution was to notice that the pattern was a fibonacci number...
// I went all the way to combinatorics to be able to get the permutations count

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// fmt.Println(climbStairs(1))
	// fmt.Println(climbStairs(2))
	// fmt.Println(climbStairs(3))
	// fmt.Println(climbStairs(5))
	// fmt.Println(climbStairs(9))
	fmt.Println(climbStairs(35))
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	ones := n
	twos := 0

	res := 1
	for ones >= 0 {
		ones -= 2
		twos++
		divisor := bigFact(ones + twos)
		dividend := bigFact(ones)
		dividend = dividend.Mul(dividend, bigFact(twos))
		permutations := divisor.Div(divisor, dividend)
		res += int(permutations.Int64())
	}
	return res
}

func bigFact(n int) *big.Int {
	res := new(big.Int).SetInt64(1)

	for i := 2; i <= n; i++ {
		res.Mul(res, new(big.Int).SetInt64(int64(i)))
	}

	return res
}
