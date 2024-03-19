// https://leetcode.com/problems/task-scheduler/description/

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	// ABCDABCDABCE
	// tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'E'}
	tasks := []byte{'A', 'B', 'A'}
	n := 2

	res := leastInterval(tasks, n)
	fmt.Println(res)
}

func leastInterval(tasks []byte, n int) int {
	// initialize maps
	tasksMap := map[byte]int{}
	tasksDistanceMap := map[byte]int{}
	for _, task := range tasks {
		_, found := tasksMap[task]
		if !found {
			tasksMap[task] = 0
		}
		tasksMap[task]++
	}

	for task := range tasksMap {
		tasksDistanceMap[task] = -1
	}

	cur := tasks[0]
	total := 0

	for len(tasksMap) > 0 {
		// check for idle
		stack := make([][]int, len(tasksMap))
		i := 0
		for task, count := range tasksMap {
			stack[i] = []int{int(task), count, tasksDistanceMap[task]}
			i++
		}
		// sort by count DESC
		slices.SortFunc(stack, func(a, b []int) int {
			return cmp.Compare(b[1], a[1])
		})
		cur = 0
		// stack[i] -> [task, count, dist]
		for _, vals := range stack {
			if vals[2] > n || vals[2] == -1 {
				cur = byte(vals[0])
				break
			}
		}

		// if cur == 0 { //idle needed
		// fmt.Println("IDLE")
		// } else {
		if cur != 0 {
			// fmt.Println(string(cur), tasksDistanceMap)
			tasksMap[cur]--
			tasksDistanceMap[cur] = 0

			if tasksMap[cur] == 0 {
				// fmt.Println("DEL", string(cur))
				delete(tasksMap, cur)
				delete(tasksDistanceMap, cur)
			}
		}

		// increase all distances
		total++
		for task := range tasksDistanceMap {
			if tasksDistanceMap[task] != -1 {
				tasksDistanceMap[task]++
			}
		}
	}

	return total
}
