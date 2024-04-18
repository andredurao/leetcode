//

package main

import "fmt"

func main() {
	grid := [][]int{{1, 1}}
	res := islandPerimeter(grid)
	fmt.Println(res)
}

func islandPerimeter(grid [][]int) int {
	total := 0
	connsMap := map[int]int{0: 4, 1: 3, 2: 2, 3: 1, 4: 0}
	rows := len(grid)
	cols := len(grid[0])

	// Up Right Down Left
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				conns := 0
				for _, dir := range dirs {
					ii := i + dir[0]
					jj := j + dir[1]
					if ii >= 0 && ii < rows &&
						jj >= 0 && jj < cols &&
						grid[ii][jj] == 1 {
						conns++
					}
				}
				total += connsMap[conns]
			}
		}
	}
	return total
}
