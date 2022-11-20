# https://leetcode.com/problems/pascals-triangle-ii/

# @param {Integer} row_index
# @return {Integer[]}
def get_row(row_index)
  result = []
  0.upto(row_index) do |i|
    if i == 0 || i == row_index
      result << 1
    else
      result << result.last * (row_index + 1 - i) / i
    end
  end
  result
end

p get_row(0)
p get_row(5)
