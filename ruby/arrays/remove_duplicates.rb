# @param {Integer[]} nums
# @return {Integer}
# Complexity analysis: Runtime O(n) Space O(n)
def remove_duplicates(nums)
  result = []

  previous = nil
  nums.each do |num|
    if !previous || (previous && num != previous)
      result << num
    end
    previous = num
  end

  index = -1
  nums.map! do |num|
    index += 1
    result[index]
  end

  result.size
end

nums = [0,0,1,1,1,2,2,3,3,4]
r = remove_duplicates(nums)
p nums, r
