# https://leetcode.com/problems/arithmetic-slices-ii-subsequence/

# @param {Integer[]} nums
# @return {Integer}
def number_of_arithmetic_slices(nums)
  # using official solution to check for timeouts
  total = 0
  caches = Array.new(nums.size){Hash.new(0)}
  0.upto(nums.size - 1) do |i|
    0.upto(i - 1) do |j|
      diff = nums[i] - nums[j]
      caches[i][diff] = caches[i][diff] + caches[j][diff] + 1
      total += caches[j][diff]
    end
  end
  total
end

nums = [2,4,6,8,10]
p number_of_arithmetic_slices(nums)

nums = [7,7,7,7,7]
p number_of_arithmetic_slices(nums)

# nums = [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]
# p number_of_arithmetic_slices(nums)
