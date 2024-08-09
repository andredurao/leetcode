# https://leetcode.com/problems/magic-squares-in-grid/

# @param {Integer[][]} grid
# @return {Integer}
def num_magic_squares_inside(grid)
  p grid
  rows = grid.size
  cols = grid[0].size
  qty = (rows - 2) * (cols - 2)
  return 0 if qty <= 0
  total = 0
  0.upto(rows-3) do |row|
    0.upto(cols-3) do |col|
      total += 1 if magic_square?(grid, row, col)
    end
  end
  total
end

# @param {Integer[][]} grid
# @param {Integer} row
# @param {Integer} col
# @return {TrueFalseClass}
def magic_square?(grid, row, col)
  sum = grid[row][col...(col+3)].sum
  # distinct check
  h = {}
  row.upto(row+2) do |i|
    col.upto(col+2) do |j|
      return false if h[grid[i][j]]
      h[grid[i][j]] = true
    end
  end

  # checksum rows
  (row+1).upto(row+2) do |i|
    return false if grid[i][col...(col+3)].sum != sum
  end
  # checksum cols
  (col).upto(col+2) do |j|
    return false if (grid[row][j] + grid[row+1][j] + grid[row+2][j]) != sum
  end
  # checksum diagonals
  (grid[row][col] + grid[row+1][col+1] + grid[row+2][col+2]) == sum &&
  (grid[row+2][col] + grid[row+1][col+1] + grid[row][col+2]) == sum
end

# ----------------------------------------

grid = [
  [4,3,8,4],
  [9,5,1,9],
  [2,7,6,2]
]
puts num_magic_squares_inside(grid)

grid = [[8]]
puts num_magic_squares_inside(grid)

grid = [[5,5,5],[5,5,5],[5,5,5]]
puts num_magic_squares_inside(grid)

grid = [[1,1,1],[1,1,1],[1,1,1]]
puts num_magic_squares_inside(grid)

# grid = [[7,0,5],[2,4,6],[3,8,1]]
# puts num_magic_squares_inside(grid)

# grid = [[10,3,5],[1,6,11],[7,9,2]]
# puts num_magic_squares_inside(grid)


# grid = [[9,9,5,1,9,5,5,7,2,5],[9,1,8,3,4,6,7,2,8,9],[4,1,1,5,9,1,5,9,6,4],[5,5,6,7,2,8,3,4,0,6],[1,9,1,8,3,1,4,2,9,4],[2,8,6,4,2,7,3,2,7,6],[9,2,5,0,7,8,2,9,5,1],[2,1,4,4,7,6,2,4,3,8],[1,2,5,3,0,5,10,8,5,2],[6,9,6,8,8,4,3,6,0,9]]
# puts num_magic_squares_inside(grid)

# grid = [[4, 3, 8, 4],[9, 5, 1, 9],[2, 7, 6, 2],[4, 3, 8, 1],[1, 6, 7, 5]]
# puts num_magic_squares_inside(grid)

# grid = [ [4, 3, 8, 4, 3],[9, 5, 1, 9, 5], [2, 7, 6, 2, 7],[4, 3, 8, 4, 3],[9, 5, 1, 9, 5],[2, 7, 6, 2, 7]]
# puts num_magic_squares_inside(grid)

# grid = [[4, 3, 8, 1, 6, 7],[9, 5, 1, 7, 8, 9],[2, 7, 6, 4, 9, 1],[8, 1, 6, 4, 3, 8],[3, 7, 9, 2, 8, 1],[4, 3, 8, 1, 6, 7],[9, 5, 1, 7, 8, 9],[2, 7, 6, 4, 9, 1],[8, 1, 6, 4, 3, 8],[1, 6, 7, 8, 9, 2]]
# puts num_magic_squares_inside(grid)

# grid = [[4, 3, 8, 4, 3, 8],[9, 5, 1, 9, 5, 1],[2, 7, 6, 2, 7, 6],[4, 3, 8, 4, 3, 8],[9, 5, 1, 9, 5, 1],[2, 7, 6, 2, 7, 6],[4, 3, 8, 4, 3, 8],[9, 5, 1, 9, 5, 1],[2, 7, 6, 2, 7, 6]]
# puts num_magic_squares_inside(grid)

# grid = [[4, 3, 8, 4, 3, 8, 4, 3, 8],[9, 5, 1, 9, 5, 1, 9, 5, 1],[2, 7, 6, 2, 7, 6, 2, 7, 6],[8, 1, 6, 4, 3, 8, 4, 3, 8],[3, 7, 9, 2, 8, 1, 9, 5, 1],[4, 3, 8, 1, 6, 7, 2, 7, 6],[9, 5, 1, 7, 8, 9, 8, 1, 6],[2, 7, 6, 4, 9, 1, 3, 7, 9],[8, 1, 6, 4, 3, 8, 4, 3, 8],[3, 7, 9, 2, 8, 1, 9, 5, 1]]
# puts num_magic_squares_inside(grid)
