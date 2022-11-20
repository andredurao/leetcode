# https://leetcode.com/problems/pascals-triangle/

# @param {Integer} num_rows
# @return {Integer[][]}
def generate(num_rows)
  rows = []
  0.upto(num_rows - 1) do |j|
    row = []
    0.upto(j) do |i|
      if i == 0 || i == j
        row << 1
      else
        row << rows[j-1][i-1] + rows[j-1][i]
      end
    end
    rows << row
  end
  rows
end


p generate(5)
