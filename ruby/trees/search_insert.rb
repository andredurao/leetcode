# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer}
def search_insert(nums, target)
  l = 0
  r = nums.size - 1
  while l <= r
    index = (l + r) / 2
    p [l,r,index,nums[index]] ; sleep 0.3
    return index if nums[index] == target
    if nums[index] > target
      r = index - 1
    else
      l = index + 1
    end
  end
  p [l,r,index,nums[index]] ; sleep 0.3

  index = (l + r) / 2
  index = 0 if index < 0
  nums[index] < target ? index+1 : index
end

nums = [1,3]
target = 0
# 0,1->1
p search_insert(nums, target)

# nums = [1,3,5,6]
# target = 5
# # 0,3->3 2,3->5 found
# p search_insert(nums, target)

# nums = [1,3,5,6]
# target = 2
# # 0,3->3 0,0->X
# p search_insert(nums, target)

# nums = [1,3,5,6]
# target = 7
# # 0,3->3 2,3->5 3,3->X
# p search_insert(nums, target)

# nums = [1,3,5,6]
# target = 6
# # 0,3->3 2,3->5 3,3
# p search_insert(nums, target)
