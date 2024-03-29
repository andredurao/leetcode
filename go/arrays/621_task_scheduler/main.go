// https://leetcode.com/problems/task-scheduler/description/

package main

import (
	"fmt"
)

func main() {
	// ABCDABCDABCE
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'E'}
	// tasks := []byte{'A', 'B', 'A'}
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

	maxCount := 0
	cur := tasks[0]
	total := 0

	for len(tasksMap) > 0 {
		// check for maxCount
		maxCount = 0
		cur = 0
		for task, count := range tasksMap {
			_, found := tasksDistanceMap[task]
			if !found {
				tasksDistanceMap[task] = -1
			}
			if count > maxCount && (tasksDistanceMap[task] > n || tasksDistanceMap[task] == -1) {
				maxCount = count
				cur = task
			}
		}

		// if cur == 0 { //idle needed
		// 	fmt.Println("IDLE")
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
