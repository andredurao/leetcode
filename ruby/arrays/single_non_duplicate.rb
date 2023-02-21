# https://leetcode.com/problems/single-element-in-a-sorted-array/

def duplicate?(index, nums)
  l = index > 0 ? nums[index-1] : nil
  r = index < nums.size - 1 ? nums[index+1] : nil
  nums[index] == l || nums[index] == r
end

# true when non dup is on the left
def shifted_right?(index, nums)
  return false if index == 0
  return false if nums[index+1] == nums[index]
  true
end

# @param {Integer[]} nums
# @return {Integer}
def single_non_duplicate(nums)
  return nums[0] if nums.size == 1
  l = 0
  r = nums.size / 2

  mid = (l + r) / 2
  while duplicate?(mid * 2, nums)
    if shifted_right?(mid * 2, nums)
      r = mid - 1
    else
      l = mid + 1
    end
    mid = (l + r) / 2
  end
  nums[mid*2]
end

# 1. non dup number will always be located in 2n index
# 2. bsearch looking up for shifted positions by the non dup
# 3. next index will always be 2n

nums = [1,1,2,3,3,4,4,8,8]
p single_non_duplicate(nums)

nums = [3,3,7,7,10,11,11]
# nums = [0,0,1,1,2, 2, 3]
p single_non_duplicate(nums)

nums = [0,0,1,1,2, 2, 3]
p single_non_duplicate(nums)
