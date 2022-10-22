# @param {Integer[]} nums
# @return {Integer[][]}
def permute(nums)
  result = []
  permute_helper(nums, result)
  result
end

def permute_helper(nums, result, repo=[])
  result << repo if nums.empty?

  nums.each_with_index do |num, index|
    permute_helper(
      nums[0...index] + nums[index+1..-1],
      result,
      repo + [num]
    )
  end
  result
end

nums = [1,2,3]
p permute(nums)
