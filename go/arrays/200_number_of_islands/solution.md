# Intuition
Initially I thought that I could iterate the rows and cols and use floodfill when land (`'1'`) was found. Then I could add up the number of islands counter. With that I could store the index of the coordinates that I've found so far in a hashmap. Before implementing I've changed my mind and decided to simply erase the islands in floodfill, so next iteration wouldn't find the island and remove the need of the hashmap.

# Approach
1. Initialize an integer total var = 0
2. Iterate the rows and cols of the grid and execute the following if the cell has land on it
2.1. Increment the total var
2.2. Call flood fill with a reference of the grid and the position on which land was found `i, j`
2.3. `[flood fill]` Check if `i` and `j` are valid positions and if the cell in that position is really `1`, otherwise return at this point
2.4. `[flood fill]` Update `grid[i][j]` to `0`
2.5. `[flood fill]` Call `2.3.` with the 4 direction neighbours

# Complexity
- Time complexity: $$O(m*n)$$

- Space complexity: $$O(m*n)$$

# Code
```
func numIslands(grid [][]byte) int {
    total := 0

    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == '1' {
                total++
                floodFill(&grid, i, j)
            }
        }
    }

    return total
}

func floodFill(grid *[][]byte, i int, j int) {
    DIRS := [][]int{{-1,0}, {0,1}, {1,0}, {0,-1}} //up right down left
    if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) || (*grid)[i][j] == '0' {
        return
    }
    (*grid)[i][j] = '0'
    for _, dir := range DIRS {
        floodFill(grid, i+dir[0], j+dir[1])
    }
}
```
