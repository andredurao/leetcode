# @param {Integer[]} nums
# @return {Integer}
def find_numbers(nums)
  total = 0
  nums.each{|num| total += 1 if num.to_s.size.even?}
  total
end

nums = [12,345,2,6,7896]
puts find_numbers(nums)
