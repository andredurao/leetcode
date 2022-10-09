# @param {Integer[]} nums
# @param {Integer} val
# @return {Integer}
# Complexity analysis: Runtime O(n), Space O(n)
def remove_element(nums, val)
  index = -1
  result = nums.reject{|num| num == val}

  nums.map! do |num|
    index += 1
    result[index]
  end
  result.length
end


# [3 2 2 3], 2
# result = [3 3], nums = [3 3 n n] 2
# [2 2 2 3], 2
# result = [3], nums = [3 n n n] 1
