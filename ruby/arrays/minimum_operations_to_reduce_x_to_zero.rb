# https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/

# @param {Integer[]} nums
# @param {Integer} x
# @return {Integer}
def min_operations(nums, x)
  total = nums.sum
  max = -1
  l = 0
  cur = 0

  nums.each_index do |r|
    cur += nums[r]

    while cur > total - x && l <= r
      cur -= nums[l]
      l += 1
    end

    if cur == total - x
      max = [max, r - l + 1].max
    end
  end

  max == -1 ? max : (nums.size - max)
end

nums = [1,1,4,2,3]
x = 5
min_operations(nums, x)
