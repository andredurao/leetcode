# @param {Integer[]} nums
# @return {Integer}
def missing_number(nums)
  # (Array(nums.min..nums.max) - nums).first
  max = nums.first
  map = {}
  nums.each do |num|
    map[num] = true
  end
  0.upto(nums.size + 1) do |val|
    return val if !map[val]
  end
end


nums = [3,0,1]
p missing_number(nums)

nums = [0,1]
p missing_number(nums)

nums = [9,6,4,2,3,5,7,0,1]
p missing_number(nums)

nums = [1]
p missing_number(nums)
