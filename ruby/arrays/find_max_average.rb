# @param {Integer[]} nums
# @param {Integer} k
# @return {Float}
def find_max_average_o_n2(nums, k)
  max_average = nil
  (0..(nums.length - k)).each do |left|
    average = nums[left...(left+k)].reduce(:+) / k.to_f
    if !max_average || average > max_average
      max_average = average
    end
  end
  max_average
end

def find_max_average_o_n(nums, k)
  max_average = nil
  last = nil
  sum = 0
  l = 0
  nums.each_with_index do |num, r|
    sum += num
    if r >= k - 1
      average = sum / k.to_f
      if !max_average || average > max_average
        max_average = average
      end
      sum -= nums[l]
      l += 1
    end

    last = num
  end
  max_average
end

nums = [1,12,-5,-6,50,3]
k = 4
p nums
puts '----------'
p find_max_average_o_n(nums, k)
