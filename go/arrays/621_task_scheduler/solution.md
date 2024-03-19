# Intuition
My first thought is that I've had to store the frequency of tasks in a hash and iterate this hash in a way that I don't need to repeat the recent `n` values.


# Approach
1. Iterate tasks and store the frequency of tasks in a hash `tasksMap`
2. Initialize an empty hash map to store the distance of tasks `tasksDistanceMap`
3. While `tasksMap` is not empty do the following:
2.1. Search for the task (`cur`) with max count and which the distance is greater than `n`, if none satisfy this condition stay idle for this cycle
2.2. Reduce the count of that task in `tasksMap`
2.3. Remove the entries in both hashes if count is zero
2.4. Increment all distances in `tasksDistanceMap`

# Complexity
- Time complexity: $$O(n^2)$$

- Space complexity: $$O(n^2)$$

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
```
