# @param {Integer[]} nums
# @return {Integer}
def largest_unique_number(nums)
  map = nums.tally
  map.select!{|item, count| count == 1}
  map.keys.max || -1
end
