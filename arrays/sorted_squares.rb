# @param {Integer[]} nums
# @return {Integer[]}
def sorted_squares(nums)
  # populate square arrays
  # O(n) time, O(n) space
  positive_squares = []
  negative_squares = []
  nums.each do |num|
    square = num * num
    if num < 0
      negative_squares << square
    else
      positive_squares << square
    end
  end

  # sorting by bubbling
  # O(n) time, O(n) space
  result = []
  while positive_squares.any? || negative_squares.any?
    #draining empty queue cases O(1)
    if negative_squares.empty?
      result += positive_squares
      positive_squares = []
    end
    if positive_squares.empty?
      result += negative_squares.reverse
      negative_squares = []
    end

    # inserting and ordering cases
    if positive_squares.any? && negative_squares.any?
      if positive_squares.first <= negative_squares.last
        result << positive_squares.shift
      else
        result << negative_squares.pop
      end
    end
  end
  result
end

nums = [-4,-1,0,3,10]
puts sorted_squares(nums)
