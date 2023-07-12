// https://leetcode.com/problems/maximum-number-of-coins-you-can-get/description/

package main

import (
	"fmt"
	"sort"
)

func main() {
	piles := []int{9, 8, 7, 6, 5, 1, 2, 3, 4}
	res := maxCoins(piles)
	fmt.Println(res)
}

func maxCoins(piles []int) int {
	sort.Slice(piles, func(i, j int) bool { return piles[i] > piles[j] })

	total := 0

	for len(piles) > 0 {
		// fmt.Println(piles)
		total += piles[1]
		// fmt.Println(total, piles[1])
		piles = piles[2 : len(piles)-1]
	}

	return total
}

// alternative
// func maxCoins(piles []int) int {
// 	sort.Slice(piles, func(i, j int) bool { return piles[i] > piles[j] })

// 	total := 0
// 	r := len(piles) - 1
// 	l := 1
// 	for {
// 		total += piles[l]
// 		l += 2
// 		r -= 1
// 		if r-l < 0 {
// 			return total
// 		}
// 	}
// }
