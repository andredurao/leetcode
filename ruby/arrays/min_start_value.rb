# @param {Integer[]} nums
# @return {Integer}
def min_start_value(nums)
  has_negative = false
  curr = 1
  smaller = 0
  transformed = nums.map do |num|
    curr += num
    has_negative = true if curr <= 0
    smaller = [smaller, curr].min
    curr
  end
  p transformed
  if has_negative
    1 -(smaller) + 1
  else
    1
  end
end

# nums = [-3,2,-3,4,2]
# p min_start_value(nums)
# nums = [1, 2]
# p min_start_value(nums)
# nums = [1,-2,-3]
# p min_start_value(nums)
nums = [5,4,-5,-5,0]
p min_start_value(nums)
