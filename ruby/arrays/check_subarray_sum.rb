# @param {Integer[]} nums
# @param {Integer} k
# @return {Boolean}
def check_subarray_sum(nums, k)
  return false if nums.size < 2
  curr_size = nums.size
  while curr_size > 1
    diff = nums.size - curr_size
    0.upto(diff) do |delta|
      sub = nums[delta..(delta + curr_size - 1)]
      p sub
      return true if sub.sum % k == 0
    end
    curr_size -= 1
  end
  false
end

# nums = [23,2,6,4,7]
# k = 13
# p check_subarray_sum(nums, k)

# nums = [23,2,4,6,6]
# k = 7
# p check_subarray_sum(nums, k)

nums = [0]
k = 1
p check_subarray_sum(nums, k)
