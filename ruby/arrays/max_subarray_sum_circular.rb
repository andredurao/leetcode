# https://leetcode.com/problems/maximum-sum-circular-subarray/

# @param {Integer[]} nums
# @return {Integer}
# def max_subarray_sum_circular(nums)
#   # from official solution (couldn't get it by myself yet)
#   cur_max = 0
#   cur_min = 0
#   sum = 0
#   max_sum = nums[0]
#   min_sum = nums[0]

#   nums.each do |num|
#     cur_max = [cur_max, 0].max + num
#     max_sum = [max_sum, cur_max].max
#     cur_min = [cur_min, 0].min + num
#     min_sum = [min_sum, cur_min].min
#     sum += num
#   end
#   sum == min_sum ? max_sum : [max_sum, sum - min_sum].max
# end


def max_subarray_sum_circular(nums)
  right_max = []
  right_max[nums.size - 1] = nums[nums.size- 1]
  suffix_sum = nums[nums.size- 1]
  i = nums.size - 2
  while i >= 0
    suffix_sum += nums[i]
    right_max[i] = [right_max[i + 1], suffix_sum].max
    i -= 1
  end
  max_sum = nums[0]
  special_sum = nums[0]
  i = 0
  prefix_sum = 0
  cur_max = 0
  while i < nums.size
    cur_max = [cur_max, 0].max + nums[i]
    #Kadane's algorithm from official solution
    max_sum = [max_sum, cur_max].max
    prefix_sum += nums[i]
    if i + 1 < nums.size
      special_sum = [special_sum, prefix_sum + right_max[i + 1]].max
    end
    i += 1
  end
  [max_sum, special_sum].max
end


nums = [1,-2,3,-2]
p max_subarray_sum_circular(nums)
