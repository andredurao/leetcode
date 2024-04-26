# Intuition
After reading the problem, I've had in mind that dynamic programming could be a way to solve this. The reason is that I've solved some of the Euler Projects problem in the past and my first contact with Dynamic Programming was with the problem 15 called : "Lattice paths" (https://projecteuler.net/problem=15), which had some similarities with this problem.

# Approach
1. Initialize a memoization matrix (`dp`) with the same dimensions as the grid and the first row having the same values as the grid
2. Iterate the grid rows, starting from the second row, while adding and memoizing the minimum path for each position, the cell's value must be added with the minimum value from the previous `dp` row, unless if that value came from the same column, an exception that occurs once per row. To deal with that exception I've decided to cache the minimum row value and it's column index from the result of a function `minValue`, which is called twice by each row iteration, 1 before the iteration starts and 2 when the index of the minimum value.
3. Return the minimum value in the last row of `dp`

* Note: I've realized later that I should have changed the follow condition from `if i == ignore` to `if i == ignore && len(*row) > 1`. The reason is that this could return a wrong result in case the input was a had more than 1 line and a single column.

# Complexity
- Time complexity: $$O(n^2)$$

- Space complexity: $$O(m*n)$$, the size of the memoization grid

# Code
```

import "math"
func minFallingPathSum(grid [][]int) int {
    dp := [][]int{}
    for row := range grid {
        if row == 0 {
            dp = append(dp, grid[0])
        } else {
            dp = append(dp, make([]int, len(grid[row])))
        }
    }


    minPrev, minPrevIndex := minValue(&dp[0], -1)
    if len(grid) == 1 {
        return minPrev
    }

    for i := 1 ; i < len(grid) ; i++ {
        for j := range grid[1] {
            if minPrevIndex == j {
                cMinValue, _ := minValue(&dp[i-1], j)
                dp[i][j] = grid[i][j] + cMinValue
            } else {
                dp[i][j] = grid[i][j] + minPrev
            }
        }
        minPrev, minPrevIndex = minValue(&dp[i], -1)
    }

    res, _ := minValue(&dp[len(dp)-1], -1)

    return res
}

func minValue(row *[]int, ignore int) (int, int) {
    res := math.MaxInt32
    index := -1
    for i, val := range *row {
        if i == ignore {
            continue
        }
        if val < res {
            res = val
            index = i
        }
    }

    return res, index
}
```
