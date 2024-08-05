# https://leetcode.com/problems/kth-distinct-string-in-an-array

# @param {String[]} arr
# @param {Integer} k
# @return {String}
def kth_distinct(arr, k)
  f = arr.tally
  arr.each do |str|
    k -= 1 if f[str] == 1
    return str if k == 0
  end

  return ""
end

# -----------------------------------

arr = ["d","b","c","b","c","a"]
k = 2
arr = ["aaa","aa","a"]
k = 1
res = kth_distinct(arr, k)
puts res
