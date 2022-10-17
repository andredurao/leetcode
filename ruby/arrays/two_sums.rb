# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer[]}
def two_sum(nums, target)
  h = {}
  nums.each_with_index do |num, index|
    remainder = target - num
    return [index, h[num]] if h[num]
    h[remainder] = index
  end
  h
end

nums = [2,7,11,15]
target = 9
p two_sum(nums, target)


nums = [3,2,4]
target = 6
p two_sum(nums, target)


nums = [3,3]
target = 6
p two_sum(nums, target)
