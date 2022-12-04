# https://leetcode.com/problems/minimum-average-difference/description/

# @param {Integer[]} nums
# @return {Integer}
def minimum_average_difference(nums)
  return 0 if nums.empty? || nums.size == 1

  result = min = -1
  sum_r = nums.sum
  sum_l = 0
  nums.each_with_index do |num, i|
    sum_l += num
    sum_r -= num
    avg_l = sum_l / (i + 1)
    avg_r = i == (nums.size - 1) ? 0 : sum_r / (nums.size - i - 1)
    current = (avg_l - avg_r).abs
    if current < min || min == -1
      result = i
      min = current
    end
  end
  result
end

nums = [0,1,0,1,0,1]
p minimum_average_difference(nums)

nums = [0,0,0,0,0]
p minimum_average_difference(nums)

nums = [2,5,3,9,5,3]
p minimum_average_difference(nums)

nums = [0]
p minimum_average_difference(nums)

nums = [1]
p minimum_average_difference(nums)

nums = []
p minimum_average_difference(nums)
