// https://leetcode.com/problems/find-center-of-star-graph/description/

package main

func main() {
	edges := [][]int{{1, 2}, {2, 3}, {4, 2}}
	println(findCenter(edges))
}

func findCenter(edges [][]int) int {
	f := map[int]int{}
	for i := range edges {
		f[edges[i][0]]++
		f[edges[i][1]]++
	}

	center := -1
	centerCount := -1
	for edge, count := range f {
		if count > centerCount {
			centerCount = count
			center = edge
		}
	}

	return center
}
