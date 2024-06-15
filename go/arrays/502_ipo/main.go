// https://leetcode.com/problems/ipo/

package main

import (
	"fmt"
	"sort"
)

func main() {
	k := 1
	w := 0
	profits := []int{1, 2, 3}
	capital := []int{0, 1, 2}
	res := findMaximizedCapital(k, w, profits, capital)
	fmt.Println(res)
}

type MinHeap []int

func (h *MinHeap) Push(x int) {
	*h = append(*h, x)
}

func (h *MinHeap) Pop() int {
	heap := *h
	n := len(heap)
	x := heap[n-1]
	*h = heap[0 : n-1]

	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	// based on editorial solution
	projects := [][]int{}
	for i := range capital {
		projects = append(projects, []int{capital[i], profits[i]})
	}
	sort.Slice(projects, func(i, j int) bool { return projects[i][0] < projects[j][0] })

	q := MinHeap{}
	index := 0
	for i := 0; i < k; i++ {
		for index < len(capital) && projects[index][0] <= w {
			// push negated
			q.Push(projects[index][1] * -1)
			index++
		}
		if len(q) == 0 {
			break
		}
		w += -q.Pop()
	}

	return w
}
