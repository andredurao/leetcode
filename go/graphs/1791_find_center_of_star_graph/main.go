// https://leetcode.com/problems/find-center-of-star-graph/description/

package main

func main() {
	edges := [][]int{{1, 2}, {2, 3}, {4, 2}}
	println(findCenter(edges))
}

func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}
	return edges[0][1]
}
