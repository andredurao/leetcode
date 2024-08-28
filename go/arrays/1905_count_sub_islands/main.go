// https://leetcode.com/problems/count-sub-islands/

package main

import "fmt"

func main() {
	// grid1 := [][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}}
	// grid2 := [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}

	grid1 := [][]int{{1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {1, 0, 1, 0, 1}}
	grid2 := [][]int{{0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {0, 1, 0, 1, 0}, {0, 1, 0, 1, 0}, {1, 0, 0, 0, 1}}
	res := countSubIslands(grid1, grid2)
	fmt.Println(res)
}

// --------------------------------

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	islands1 := getIslands(&grid1)
	islands2 := getIslands(&grid2)

	// invert islands1: from (island # -> []pixels) to (pixel -> island #)
	reverseIsland1 := map[int]int{}
	for land, pixels := range islands1 {
		for pixel := range pixels {
			reverseIsland1[pixel] = land
		}
	}
	// fmt.Println(islands1)
	// fmt.Println(islands2)
	// fmt.Println(reverseIsland1)

	// check for sub islands
	total := 0
	for _, pixels := range islands2 {
		currentIsland := -1
		subIsland := true
		for pixel := range pixels {
			island, found := reverseIsland1[pixel]
			if !found {
				subIsland = false
				continue
			}
			if currentIsland == -1 {
				currentIsland = island
			} else if currentIsland != island {
				subIsland = false
				continue
			}
		}
		if subIsland {
			total++
		}
	}

	return total
}

func getIslands(grid *[][]int) map[int]map[int]struct{} {
	islands := map[int]map[int]struct{}{}
	islandCount := 2
	for i := range *grid {
		for j := range (*grid)[i] {
			if (*grid)[i][j] == 1 {
				land := map[int]struct{}{}
				floodfill(grid, land, islandCount, i, j, len((*grid)[0]))
				islands[islandCount] = land
				islandCount++
			}
		}
	}
	return islands
}

func floodfill(grid *[][]int, land map[int]struct{}, count, i, j, width int) {
	(*grid)[i][j] = count
	land[pos2int(width, i, j)] = struct{}{}

	positions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, pos := range positions {
		ii := i + pos[0]
		jj := j + pos[1]
		if ii >= 0 && jj >= 0 && ii < len(*grid) && jj < len((*grid)[0]) && (*grid)[ii][jj] == 1 {
			floodfill(grid, land, count, ii, jj, width)
		}
	}
}

func pos2int(width, i, j int) int {
	return width*i + j
}
