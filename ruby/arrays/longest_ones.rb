# https://leetcode.com/problems/max-consecutive-ones-iii/
# @param {Integer[]} nums
# @param {Integer} k
# @return {Integer}
def longest_ones(nums, k)
  left = 0
  zeroes = k
  nums.each_with_index do |num, index|
    zeroes += (num - 1) # decrease 1 if value = 0
    if zeroes < 0
      zeroes += 1 if nums[left] == 0
      left += 1
    end
  end
  nums.size - left
end

nums = [1,1,1,0,0,0,1,1,1,1,0]
k = 2
p longest_ones(nums, k)

nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1]
k = 3
p longest_ones(nums, k)
