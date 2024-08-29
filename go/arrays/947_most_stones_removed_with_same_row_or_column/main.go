//

package main

import (
	"fmt"
)

func main() {
	stones := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}
	res := removeStones(stones)
	fmt.Println(res)
}

// -----------------

type Stone struct {
	Row         int
	Col         int
	Connections []*Stone
	Visited     bool
}

func removeStones(stones [][]int) int {
	stonesList := make([]*Stone, len(stones))
	for i, stone := range stones {
		stonesList[i] = &Stone{Row: stone[0], Col: stone[1], Visited: false}
	}

	for i := range stones {
		for j := i + 1; j < len(stones); j++ {
			if stonesList[i].Row == stonesList[j].Row || stonesList[i].Col == stonesList[j].Col {
				stonesList[i].Connections = append(stonesList[i].Connections, stonesList[j])
				stonesList[j].Connections = append(stonesList[j].Connections, stonesList[i])
			}
		}
	}

	visitedCount := 0
	for _, stone := range stonesList {
		if !stone.Visited {
			dfs(stone, &stonesList)
			visitedCount++
		}
	}

	return len(stones) - visitedCount
}

func dfs(stone *Stone, stonesList *[]*Stone) {
	stone.Visited = true
	for _, connection := range stone.Connections {
		if !connection.Visited {
			dfs(connection, stonesList)
		}
	}
}
