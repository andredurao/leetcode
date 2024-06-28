# Intuition
I noticed that the importance was also proportical to the amount of connections that each edge had. I've decided to create a list of frequencies and their edge numbers and sort the list later to get their labels (0 -> 2, 1 -> 4, ...) after that I've iterated through the roads again and added the values

# Approach
1. Initialize a 2d array `int[][]` with each road number and the value 0
2. Iterate through roads and increment `f` at each edge used `road[0]` and `road[1]`
3. Sort `f` by their counts `f[x][1]`
4. Initialize an array `labels` of size `n` and iterate though the values of `f` assigning the values of each label on `f[i][0]` the value of `i+1`, because the importance values start at 1
5. Initialize the `total` value (return value)
6. Iterate through the values of roads one last time and sum the values of each node indexed by the array `labels` defined previously

# Complexity
- Time complexity: $$O(m * log(n))$$

- Space complexity: $$O(n)$$

# Code
```
import "sort"

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
        labels[f[i][0]] = i+1
    }

    total := int64(0)
    for _, road := range roads {
        total += int64(labels[road[0]]+labels[road[1]])
    }

    return total
}
```
