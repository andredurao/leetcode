# @param {Integer[]} arr
# @return {Boolean}
def unique_occurrences(arr)
  counts = arr.tally.values
  counts.tally.values.all?{|v| v == 1}
end

arr = [1,2,2,1,1,3]
p unique_occurrences(arr)
