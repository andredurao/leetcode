// https://leetcode.com/problems/ipo/

package main

import (
	"container/heap"
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

// check https://pkg.go.dev/container/heap
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	// based on editorial solution
	projects := [][]int{}
	for i := range capital {
		projects = append(projects, []int{capital[i], profits[i]})
	}
	sort.Slice(projects, func(i, j int) bool { return projects[i][0] < projects[j][0] })

	h := &IntHeap{}
	heap.Init(h)

	index := 0
	for i := 0; i < k; i++ {
		for index < len(capital) && projects[index][0] <= w {
			// push negated
			// q.Push(projects[index][1] * -1)
			heap.Push(h, projects[index][1]*-1)

			index++
		}
		if h.Len() == 0 {
			break
		}
		w += heap.Pop(h).(int) * -1
	}

	return w
}
