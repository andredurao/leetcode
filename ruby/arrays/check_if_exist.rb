# Given an array arr of integers, check if there exist two indices i and j such that :
# i != j
# 0 <= i, j < arr.length
# arr[i] == 2 * arr[j]

# Example 1:
# Input: arr = [10,2,5,3]
# Output: true
# Explanation: For i = 0 and j = 2, arr[i] == 10 == 2 * 5 == 2 * arr[j]

# Example 2:
# Input: arr = [3,1,7,11]
# Output: false
# Explanation: There is no i and j that satisfy the conditions.
require 'set'
# @param {Integer[]} arr
# @return {Boolean}
def check_if_exist(arr)
  map = arr.tally
  arr.each do |value|
    next if value == 0 && map[value] < 2
    if !!map[value*2] || value.even? && !!map[value/2]
      return true
    end
  end
  false
end

# arr = [10,2,5,3]
# p check_if_exist(arr)

# arr = [3,1,7,11]
# p check_if_exist(arr)

#        0  1  2  3   4 5  6
# arr = [-2 0  10 -19 4 6 -8]
#           i j
#        i = 1, j = 2
#
# arr[i] = 0
# arr[j] = 10
#
# 0 * 2  == 0 (ok)
# 10 / 2 == 5 (not ok)

arr = [-2,0,10,-19,4,6,-8]
p check_if_exist(arr)

# arr = [0, 0]
# p check_if_exist(arr)

# TODO:  answer this https://leetcode.com/explore/learn/card/fun-with-arrays/527/searching-for-items-in-an-array/3250/discuss/2399713/Remove-zeroes-for-acceptance
