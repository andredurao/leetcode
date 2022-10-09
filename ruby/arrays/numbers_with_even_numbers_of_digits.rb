# @param {Integer[]} nums
# @return {Integer}
def find_numbers(nums)
  total = 0
  nums.each{|num| total += 1 if num.to_s.size.even?}
  total
end

nums = [12,345,2,6,7896]
puts find_numbers(nums)

# Another solution would be using truncate of log(num) at base 10.
# It wound follow the amount of digits per each number
