// https://leetcode.com/problems/find-all-groups-of-farmland/

package main

import (
	"fmt"
	"math"
)

func main() {
	land := [][]int{
		{1, 0, 0},
		{0, 1, 1},
		{0, 1, 1},
	}
	res := findFarmland(land)
	fmt.Println(res)
}

func findFarmland(land [][]int) [][]int {
	res := [][]int{}
	for i := range land {
		for j := range land[i] {
			if land[i][j] == 1 {
				landGroup := []int{math.MaxInt32, math.MaxInt32, 0, 0}
				searchForBounds(&land, i, j, &landGroup)
				res = append(res, landGroup)
			}
		}
	}

	return res
}

// search for bounds using flood fill
func searchForBounds(land *[][]int, i int, j int, landGroup *[]int) {
	DIRS := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} //up right down left
	if i < 0 || j < 0 || i >= len(*land) || j >= len((*land)[0]) || (*land)[i][j] == 0 {
		return
	}
	(*land)[i][j] = 0
	(*landGroup)[0] = min((*landGroup)[0], i)
	(*landGroup)[1] = min((*landGroup)[1], j)
	(*landGroup)[2] = max((*landGroup)[2], i)
	(*landGroup)[3] = max((*landGroup)[3], j)

	for _, dir := range DIRS {
		searchForBounds(land, i+dir[0], j+dir[1], landGroup)
	}
}
