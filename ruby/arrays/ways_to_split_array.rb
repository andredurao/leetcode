# https://leetcode.com/problems/number-of-ways-to-split-array/
# @param {Integer[]} nums
# @return {Integer}
def ways_to_split_array(nums)
  sum = nums.sum
  sums = []
  result = 0
  total = 0
  nums.each do |value|
    total += value
    sums << total
  end
  index = 0
  while index < nums.size - 1
    current = sums[index]
    valid = current >= sums.last - current
    result += 1 if valid
    index += 1
  end
  result
end

nums = [10,4,-8,7]
p ways_to_split_array(nums)

p '----------------------'

nums = [2,3,1,0]
p ways_to_split_array(nums)
