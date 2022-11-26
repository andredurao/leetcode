# https://leetcode.com/problems/word-search/

require 'set'

# @param {Character[][]} board
# @param {String} word
# @return {Boolean}
def exist(board, word)
  board_chars = Set.new
  board.each do |row|
    row.each do |char|
      board_chars.add(char)
    end
  end
  return false if (Set.new(word.chars) - board_chars).any?

  height = board.size
  width = board[0].size
  chars = word.chars

  # p '[pos, chars, board, index, height, width]'
  board.each_with_index do |row, i|
    row.each_with_index do |char, j|
      if char == word[0]
        result = visit([i,j], chars, board, 0, height, width)
        return true if result
      end
    end
  end
  false
end

def visit(pos, chars, board, index, height, width)
  # p [pos, chars, board, index]
  return true if index == chars.size
  # check if pos is out of bounds
  return false if pos[0] < 0 || pos[0] >= height || pos[1] < 0 || pos[1] >= width

  expected = chars[index]
  return false if board[pos[0]][pos[1]] != expected
  board[pos[0]][pos[1]] = 'x' # mark visited

  [[1,0],[-1,0],[0,1],[0,-1]].each do |to|
    new_pos = [pos[0] + to[0], pos[1] + to[1]]
    result = visit(new_pos, chars, board, index + 1, height, width)
    return true if result
  end

  board[pos[0]][pos[1]] = expected # restore value
  false
end

board = [
  ["A","B","C","E"],
  ["S","F","C","S"],
  ["A","D","E","E"]
]
word = "ABCCED"
p exist(board, word)

# board = [
#   ["A","B","C","E"],
#   ["S","F","C","S"],
#   ["A","D","E","E"]
# ]
# word = "SEE"
# p exist(board, word)

# board = [
#   ["A","B","C","E"],
#   ["S","F","C","S"],
#   ["A","D","E","E"]
# ]
# word = "ABCB"
# p exist(board, word)

# board = [["a"]]
# word = "b"
# p exist(board, word)

# board = [["a", "a"]]
# word = "aaa"
# p exist(board, word)

# board = [
#   ["A","A","A","A","A","A"],
#   ["A","A","A","A","A","A"],
#   ["A","A","A","A","A","A"],
#   ["A","A","A","A","A","A"],
#   ["A","A","A","A","A","A"],
#   ["A","A","A","A","A","A"]]
# word = "AAAAAAAAAAAAAAB"
# p exist(board, word)


# board = [
#   ["A","A","A","A","A","A"],
#   ["A","A","A","A","A","A"],
#   ["A","A","A","A","A","A"],
#   ["A","A","A","A","A","A"],
#   ["A","A","A","A","A","A"],
#   ["A","A","A","A","A","A"]]
# word = "AAAAAAAAAAAAAAB"
# p exist(board, word)

board = [
  ["A","A","A","A","A","A"],
  ["A","A","A","A","A","A"],
  ["A","A","A","A","A","A"],
  ["A","A","A","A","A","A"],
  ["A","A","A","A","A","B"],
  ["A","A","A","A","B","A"]]
word = "AAAAAAAAAAAAABB"
p exist(board, word)
