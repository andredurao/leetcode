# idea that I'd like to test:
# 1: convert from binary to number
# 2: move l & r with shr & shl while number > 0
# 3: test if the number of bits is equal with popcount?


# @param {Integer[]} nums
# @return {Integer}
def find_max_length(nums)
  map = { 0 => -1 }
  max_length = sum = 0
  nums.each_with_index do |num, index|
    # sum += num == 0 ? -1 : 1
    sum += (num * 2) - 1
    if map[sum]
      curr_len = index - map[sum]
      max_length = curr_len if curr_len > max_length
    else
      map[sum] = index
    end
  end
  max_length
end

nums = [0,1]
p find_max_length(nums)
# p '------------------'
# nums = [0,1,0]
# p find_max_length(nums)
# from solution (couldn't find O(n) by myself): https://leetcode.com/problems/contiguous-array/solution/
nums = [0,1,0,0,1,1,0]
p find_max_length(nums)
