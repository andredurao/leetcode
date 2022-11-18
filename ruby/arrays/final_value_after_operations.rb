# https://leetcode.com/problems/final-value-of-variable-after-performing-operations/

# @param {String[]} operations
# @return {Integer}
def final_value_after_operations(operations)
  map = { '++X' => 1, 'X++' => 1, 'X--' => -1, '--X' => -1 }
  total = 0
  operations.each do |op|
    total += map[op]
  end
  total
end

p final_value_after_operations(["--X","X++","X++"])
