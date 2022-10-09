# @param {Integer[]} nums
# @return {Integer[]}
def sorted_squares(nums)
  l = 0
  r = nums.size - 1
  result = []

  while l != r
    if nums[l].abs >= nums[r]
      result.unshift(nums[l] ** 2)
      l += 1
    else
      result.unshift(nums[r] ** 2)
      r -= 1
    end
  end

  result.unshift(nums[r] ** 2)
end

nums = [-4,-1,0,3,10]
puts sorted_squares(nums)
