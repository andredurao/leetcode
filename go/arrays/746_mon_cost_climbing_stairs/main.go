// https://leetcode.com/problems/min-cost-climbing-stairs/description/

package main

import (
	"fmt"
	"math"
)

func main() {

	x := min(1, 2)
	fmt.Println("Hello World!", x)

	cost := []int{10, 15, 20}
	result := minCostClimbingStairs(cost)
	fmt.Println(result)
}

func minCostClimbingStairs(cost []int) int {
	minCost := make([]int, len(cost)+1)
	minCost[len(cost)] = 0

	for i := len(cost) - 1; i >= 0; i-- {
		n1Step := minCost[i+1]
		n2Step := n1Step
		if (i + 2) < len(minCost) {
			n2Step = minCost[i+2]
		}
		minCost[i] = cost[i] + int(math.Min(float64(n1Step), float64(n2Step)))

	}

	//fmt.Printf("%v\n", minCost)

	return int(math.Min(float64(minCost[0]), float64(minCost[1])))
}
