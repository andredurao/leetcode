//

package main

import "fmt"

func main() {
	grid := [][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}
	res := getMaximumGold(grid)
	fmt.Println(res)
}

func getMaximumGold(grid [][]int) int {
	total := 0
	for i := range grid {
		for j := range grid[i] {
			total = max(total, dfs(i, j, &grid))
		}
	}

	return total
}

func dfs(i int, j int, grid *[][]int) int {
	DIRS := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) || (*grid)[i][j] == 0 {
		return 0
	}

	res := 0
	tmp := (*grid)[i][j]
	(*grid)[i][j] = 0

	for _, dir := range DIRS {
		dfsRes := dfs(i+dir[0], j+dir[1], grid)
		res = max(res, dfsRes)
	}

	(*grid)[i][j] = tmp
	return res + tmp
}
