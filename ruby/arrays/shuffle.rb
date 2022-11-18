# https://leetcode.com/problems/shuffle-the-array/

# @param {Integer[]} nums
# @param {Integer} n
# @return {Integer[]}
def shuffle(nums, n)
  result = []
  0.upto(n-1) do |i|
    result << nums[i]
    result << nums[n+i]
  end
  result
end


nums = [2,5,1,3,4,7]
n = 3
p shuffle(nums, n)
