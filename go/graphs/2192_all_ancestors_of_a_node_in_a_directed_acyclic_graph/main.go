//

package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 8
	edges := [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}
	res := getAncestors(n, edges)
	fmt.Println(res)
}

func getAncestors(n int, edges [][]int) [][]int {
	ancestorsMap := []map[int]struct{}{}

	for i := 0; i < n; i++ {
		ancestorsMap = append(ancestorsMap, map[int]struct{}{})
	}

	for _, edge := range edges {
		ancestorsMap[edge[1]][edge[0]] = struct{}{}
	}

	ancestors := make([][]int, n)

	for i, curMap := range ancestorsMap {
		keys := make([]int, len(curMap))
		j := 0
		for k := range curMap {
			keys[j] = k
			j++
		}
		sort.IntSlice.Sort(keys)
		ancestors[i] = keys
	}

	return ancestors
}
