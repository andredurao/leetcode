// https://leetcode.com/problems/number-of-islands/

package main

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	res := numIslands(grid)
	println(res)
}

func numIslands(grid [][]byte) int {
	total := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				total++
				floodFill(&grid, i, j)
			}
		}
	}

	return total
}

func floodFill(grid *[][]byte, i int, j int) {
	DIRS := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} //up right down left
	if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) || (*grid)[i][j] == '0' {
		return
	}
	(*grid)[i][j] = '0'
	for _, dir := range DIRS {
		floodFill(grid, i+dir[0], j+dir[1])
	}
}
