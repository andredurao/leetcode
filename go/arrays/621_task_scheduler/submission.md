# Intuition
My first thoughts where that I've had to store the frequency of tasks in a hash and iterate this hash in a way that I don't need to repeat the recent `n` values.

# Approach
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```
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
```
