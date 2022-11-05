# @param {Integer[]} nums
# @return {Integer[]}
def running_sum(nums)
  buff = 0
  nums.map do |num|
    buff += num
    buff
  end
end

nums = [1,1,1,1,1]
p running_sum(nums)
