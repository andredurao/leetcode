# @param {Integer[]} nums
# @return {Integer}
def number_of_arithmetic_slices(nums)
  # p nums
  total = 0
  i = (1 << nums.size) - 1
  while i > 0
    bits_str = i.to_s(2)
    bits = ('0' * (nums.size - bits_str.length)) << bits_str
    slice = []
    nums.each_with_index{|num, index| slice << num if bits[index] == '1'}
    if slice.size > 2
      if sequence?(slice)
        # p bits
        # p slice
        total += 1
      end
    end
    # p slice
    i -= 1
  end
  total
end

# TODO: Cache slices and their responces in a hash
# ex: there may be some improving by caching [1,2,3] and using it while performing [1,2,3,4]?
def sequence?(slice)
  last = nil
  1.upto(slice.size - 1) do |i|
    delta = slice[i] - slice[i-1]
    last ||= delta
    return false if last != delta
    last = delta
  end
  true
end

nums = [2,4,6,8,10]
p number_of_arithmetic_slices(nums)

nums = [7,7,7,7,7]
p number_of_arithmetic_slices(nums)
