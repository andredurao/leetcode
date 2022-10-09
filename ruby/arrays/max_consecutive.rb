def find_max_consecutive_ones(nums)
  max = 0
  len = 0
  nums.each_with_index do |num|
    if num == 0
      len = 0
    else
      len += 1
    end

    max = len if len >= max
  end
  max
end

nums = [1,1,0,1,1,1]
puts find_max_consecutive_ones(nums)
