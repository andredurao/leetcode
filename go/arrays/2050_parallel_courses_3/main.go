// https://leetcode.com/problems/parallel-courses-iii/

package main

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	// n := 3
	// relations := [][]int{{1, 3}, {2, 3}}
	// time := []int{3, 2, 5}

	n := 5
	relations := [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}}
	time := []int{1, 2, 3, 4, 5}

	result := minimumTime(n, relations, time)
	fmt.Println(result)
}

// following editorial
func minimumTime(n int, relations [][]int, time []int) int {
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}

	// indegree
	indegree := make([]int, n)
	for _, edge := range relations {
		from := edge[0] - 1
		to := edge[1] - 1
		graph[from] = append(graph[from], to)
		indegree[to]++
	}

	queue := list.New()
	maxTime := make([]int, n)

	// initialize list
	for node := 0; node < n; node++ {
		if indegree[node] == 0 {
			queue.PushBack(node)
			maxTime[node] = time[node]
		}
	}

	for queue.Len() > 0 {
		head := queue.Front()
		queue.Remove(head)
		node := head.Value.(int)
		neighbours := graph[node]
		for _, neighbour := range neighbours {
			max := int(math.Max(
				float64(maxTime[neighbour]),
				float64(time[neighbour]+maxTime[node]),
			))
			maxTime[neighbour] = max
			indegree[neighbour] = indegree[neighbour] - 1
			if indegree[neighbour] == 0 {
				queue.PushBack(neighbour)
			}
		}
	}

	result := 0
	for i := 0; i < n; i++ {
		result = int(math.Max(
			float64(result),
			float64(maxTime[i]),
		))
	}

	return result
}
