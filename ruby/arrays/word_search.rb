# https://leetcode.com/problems/word-search/
require 'set'
# @param {Character[][]} board
# @param {String} word
# @return {Boolean}
def exist(board, word)
  height = board.size
  width = board[0].size
  charmap = {}
  board.each_with_index do |row, i|
    row.each_with_index do |char, j|
      charmap[char] ||= []
      charmap[char] << [i,j]
    end
  end


  chars = word.chars
  path = [chars.shift]
  (charmap[path.first] || []).each do |pos|
    result = visit(pos, path, chars, Set.new([pos]), board, height, width)
    return true if result
  end
  false
end

def visit(pos, path, chars, positions, board, height, width)
  dirs = [[1,0],[-1,0],[0,1],[0,-1]]
  return true if chars.empty?

  chars = chars.dup
  expected = chars.shift
  dirs.each do |to|
    new_pos = [pos[0] + to[0], pos[1] + to[1]]
    # p [pos, new_pos, path, chars, expected]
    if valid?(new_pos, height, width) && !positions.include?(new_pos) && board[new_pos[0]][new_pos[1]] == expected
      result = visit(new_pos, path.dup << expected, chars, positions.dup.add(new_pos), board, height, width)
      return true if result
    end
  end
  false
end

def valid?(pos, height, width)
  pos[0] >= 0 && pos[0] < height && pos[1] >= 0 && pos[1] < width
end

# board = [
#   ["A","B","C","E"],
#   ["S","F","C","S"],
#   ["A","D","E","E"]
# ]
# word = "ABCCED"
# p exist(board, word)

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

board = [
  ["A","A","A","A","A","A"],
  ["A","A","A","A","A","A"],
  ["A","A","A","A","A","A"],
  ["A","A","A","A","A","A"],
  ["A","A","A","A","A","A"],
  ["A","A","A","A","A","A"]]
word = "AAAAAAAAAAAAAAB"
p exist(board, word)

