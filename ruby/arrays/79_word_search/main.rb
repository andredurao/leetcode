# @param {Character[][]} board
# @param {String} word
# @return {Boolean}
def exist(board, word)
  # Count the frequency of repeated chars and
  char_frequency = []
  last = nil
  word.chars.each do |char|
    if char != last
      char_frequency << 1
    else
      char_frequency[char_frequency.size - 1] += 1
    end
    last = char
  end

  # Check if the frequency counts is increasing or decreasing
  theta = 0
  if char_frequency.size > 1
    1.upto(char_frequency.size - 1) do |i|
      theta += char_frequency[i-1] <=> char_frequency[i]
    end
  end

  # reverse word if the frequency is increasing
  word.reverse! if theta > 0

  @height = board.size
  @width = board[0].size
  @word = word
  @board = board

  # p '[pos, chars, board, index, height, width]'
  board.each_with_index do |row, i|
    row.each_with_index do |char, j|
      if char == word[0]
        result = visit(i, j, 0)
        return true if result
      end
    end
  end
  false
end

def visit(i, j, index)
  # p [pos, index, @height, @width]
  return true if index == @word.size
  # check if pos is out of bounds
  return false if i < 0 || i >= @height || j < 0 || j >= @width

  expected = @word[index]
  # binding.irb
  return false if @board[i][j] != expected
  @board[i][j] = 'x' # mark visited

  [[1,0],[-1,0],[0,1],[0,-1]].each do |to|
    result = visit(i + to[0], j + to[1], index + 1)
    return true if result
  end

  @board[i][j] = expected # restore value
  false
end

board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]
word = "ABCCED"
puts exist(board, word)
