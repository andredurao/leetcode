// https://leetcode.com/problems/maximum-total-importance-of-roads/description/

package main

import "sort"

func main() {
	n := 5
	roads := [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}}
	res := maximumImportance(n, roads)
	println(res)
}

func maximumImportance(n int, roads [][]int) int64 {
	f := [][]int{}
	for i := 0; i < n; i++ {
		f = append(f, []int{i, 0})
	}
	for i := range roads {
		f[roads[i][0]][1]++
		f[roads[i][1]][1]++
	}

	sort.Slice(f, func(i, j int) bool { return f[i][1] < f[j][1] })

	// defining labels
	labels := make([]int, n)
	for i := range f {
		labels[f[i][0]] = i + 1
	}

	total := int64(0)
	for _, road := range roads {
		total += int64(labels[road[0]] + labels[road[1]])
	}

	return total
}
