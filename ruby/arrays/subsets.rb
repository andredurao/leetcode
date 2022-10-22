# @param {Integer[]} nums
# @return {Integer[][]}
def subsets(nums)
  return [] if nums.empty?
  # binary collections to subsets
  result = []
  0.upto((1 << nums.size) - 1) do |num|
    bits = bits_enabled(num)
    set = bits.reduce([]) do |array, bit|
      array << nums[bit] if nums[bit]
      array
    end
    result << set
  end
  result
end

# This explanation from example 8.4 is more simpler to understand than the recursive IMO
# array of bits that are enabled on number n
def bits_enabled(n)
  result = []
  index = 0
  while n > 0
    result << index if n % 2 == 1
    n = n / 2
    index += 1
  end
  result
end

p subsets([1, 2])
p subsets([1, 2, 3])
