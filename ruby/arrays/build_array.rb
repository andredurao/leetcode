# https://leetcode.com/problems/build-array-from-permutation/

# @param {Integer[]} nums
# @return {Integer[]}
def build_array(nums)
  result = []
  nums.each_index{|i| result << nums[nums[i]]}
  result
end

p build_array([0,2,1,5,3,4])
