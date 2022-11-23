# https://leetcode.com/problems/valid-sudoku/

require 'set'

# @param {Character[][]} board
# @return {Boolean}
def is_valid_sudoku(board)
  d = board.size
  # check rows
  board.each do |row|
    s = Set.new
    row.each do |col|
      next if col == '.'
      return false if s.include?(col)
      s.add(col)
    end
  end

  # check cols
  0.upto(d - 1) do |row|
    s = Set.new
    0.upto(d - 1) do |col|
      val = board[col][row]
      next if val == '.'
      return false if s.include?(val)
      s.add(val)
    end
  end

  # check grids
  0.upto(2) do |j|
    0.upto(2) do |i|
      s = Set.new
      0.upto(2) do |row|
        0.upto(2) do |col|
          val = board[3 * j + row][3 * i + col]
          next if val == '.'
          return false if s.include?(val)
          s.add(val)
        end
      end
    end
  end

  true
end

board = [
  %w[5 3 . . 7 . . . .],
  %w[6 . . 1 9 5 . . .],
  %w[. 9 8 . . . . 6 .],
  %w[8 . . . 6 . . . 3],
  %w[4 . . 8 . 3 . . 1],
  %w[7 . . . 2 . . . 6],
  %w[. 6 . . . . 2 8 .],
  %w[. . . 4 1 9 . . 5],
  %w[. . . . 8 . . 7 9]]
p is_valid_sudoku(board)

board = [
  %w[8 3 . . 7 . . . .],
  %w[6 . . 1 9 5 . . .],
  %w[. 9 8 . . . . 6 .],
  %w[8 . . . 6 . . . 3],
  %w[4 . . 8 . 3 . . 1],
  %w[7 . . . 2 . . . 6],
  %w[. 6 . . . . 2 8 .],
  %w[. . . 4 1 9 . . 5],
  %w[. . . . 8 . . 7 9]]
p is_valid_sudoku(board)

board = [
  %w[. . 4 . . . 6 3 .],
  %w[. . . . . . . . .],
  %w[5 . . . . . . 9 .],
  %w[. . . 5 6 . . . .],
  %w[4 . 3 . . . . . 1],
  %w[. . . 7 . . . . .],
  %w[. . . 5 . . . . .],
  %w[. . . . . . . . .],
  %w[. . . . . . . . .]]
p is_valid_sudoku(board)


board = [
  %w[. . . . 5 . . 1 .],
  %w[. 4 . 3 . . . . .],
  %w[. . . . . 3 . . 1],
  %w[8 . . . . . . 2 .],
  %w[. . 2 . 7 . . . .],
  %w[. 1 5 . . . . . .],
  %w[. . . . . 2 . . .],
  %w[. 2 . 9 . . . . .],
  %w[. . 4 . . . . . .]]
p is_valid_sudoku(board)
