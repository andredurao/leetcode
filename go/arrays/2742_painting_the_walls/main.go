// https://leetcode.com/problems/painting-the-walls/description/

package main

import (
	"fmt"
	"sort"
)

func main() {
	// cost := []int{1, 2, 3, 2}
	// time := []int{1, 2, 3, 2}
	cost := []int{26, 53, 10, 24, 25, 20, 63, 51}
	time := []int{1, 1, 1, 1, 2, 2, 2, 1}
	result := paintWalls(cost, time)
	fmt.Println(result)

	cost = []int{8, 7, 5, 15}
	time = []int{1, 1, 2, 1}
	result = paintWalls(cost, time)
	fmt.Println(result)

	cost = []int{42, 8, 28, 35, 21, 13, 21, 35}
	time = []int{2, 1, 1, 1, 2, 1, 1, 2}
	result = paintWalls(cost, time)
	fmt.Println(result)
}

func paintWalls(cost []int, time []int) int {
	// build an array joining cost and time
	sortedCosts := make([][]int, len(cost))
	for i, val := range cost {
		sortedCosts[i] = []int{val, time[i]}
	}
	// sort by cost ASC, time DESC
	sort.Slice(sortedCosts, func(i, j int) bool {
		return (sortedCosts[i][0] < sortedCosts[j][0]) || (sortedCosts[i][1] > sortedCosts[j][1])
	})

	fmt.Printf("%v\n", cost)
	fmt.Printf("%v\n", sortedCosts)

	l := 0
	r := len(cost) - 1
	total := 0

	fmt.Println("l", l, "r", r)
	for l <= r {
		// choose minimum cost wall (probably on the left)
		paidIndex := minWall(sortedCosts, l, r)
		total += sortedCosts[paidIndex][0]
		// paint (time) free walls
		r = r - sortedCosts[paidIndex][1]
		l++
		fmt.Println("l", l, "r", r, "total", total)
	}

	return total
}

func minWall(sortedCosts [][]int, i int, j int) int {
	if sortedCosts[i][0] < sortedCosts[j][0] {
		return i
	} else if sortedCosts[i][0] == sortedCosts[j][0] {
		if sortedCosts[i][1] >= sortedCosts[j][1] {
			return i
		} else {
			return j
		}
	} else {
		return j
	}
}
