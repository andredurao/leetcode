//

package main

import "fmt"

func main() {
	// n := 4
	// edges := [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}
	// distanceThreshold := 4
	// res := findTheCity(n, edges, distanceThreshold)
	// fmt.Println(res)

	n := 5
	edges := [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}
	distanceThreshold := 2
	res := findTheCity(n, edges, distanceThreshold)
	fmt.Println(res)

}

// ---------------------------------------------------------------

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	neighboursMap := map[int][][]int{}
	for _, edge := range edges {
		neighboursMap[edge[0]] = append(neighboursMap[edge[0]], []int{edge[1], edge[2]})
		neighboursMap[edge[1]] = append(neighboursMap[edge[1]], []int{edge[0], edge[2]})
	}

	counts := map[int][]int{}
	minCount := 999
	for i := 0; i < n; i++ {
		visits := map[int]struct{}{}
		visit(i, 0, distanceThreshold, visits, neighboursMap)
		counts[len(visits)] = append(counts[len(visits)], i)
		minCount = min(minCount, len(visits))
	}

	res := counts[minCount][0]
	for i := range counts[minCount] {
		res = max(res, counts[minCount][i])
	}

	return res
}

func visit(node, d, distanceThreshold int, visits map[int]struct{}, neighboursMap map[int][][]int) {
	if d > distanceThreshold {
		return
	}
	if _, found := visits[node]; found {
		return
	}
	visits[node] = struct{}{}
	for _, neighbours := range neighboursMap[node] {
		visit(neighbours[0], d+neighbours[1], distanceThreshold, visits, neighboursMap)
	}
}
