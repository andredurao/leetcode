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

	cost = []int{42, 8, 28, 35, 21, 13, 21, 35}
	time = []int{2, 1, 1, 1, 2, 1, 1, 2}
	result = paintWalls(cost, time)
	fmt.Println(result)
}

func paintWalls(cost []int, time []int) int {
	memo := make([][]int, len(cost))
	for i := range cost {
		memo[i] = make([]int, len(cost)+1)
	}
	return dp(0, len(cost), cost, time, &memo)
}

func dp(i int, remain int, cost []int, time []int, memo *[][]int) int {
	// fmt.Printf("%v\n", *memo)
	// fmt.Println("dp")
	if remain <= 0 {
		return 0
	}

	if i == len(cost) {
		return 1000000000
	}

	if (*memo)[i][remain] != 0 {
		return (*memo)[i][remain]
	}

	paint := cost[i] + dp(i+1, remain-1-time[i], cost, time, memo)
	dontPaint := dp(i+1, remain, cost, time, memo)
	(*memo)[i][remain] = int(math.Min(float64(paint), float64(dontPaint)))

	return (*memo)[i][remain]
}
