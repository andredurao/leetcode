# https://leetcode.com/problems/insert-interval/

# @param {Integer[][]} intervals
# @param {Integer[]} new_interval
# @return {Integer[][]}
def insert(intervals, new_interval)
  index = intervals.bsearch_index{|interval| new_interval[0] < interval[0]} || intervals.size
  intervals.insert(index, new_interval)
  last = nil
  result = intervals
  compress(result)
end

def compress(array)
  previous = nil
  result = []
  array.each do |item|
    if previous
      if (previous[0]..previous[1]).include?(item[0])
        previous += item
        previous = [previous.min, previous.max]
      else
        result << previous
        previous = item
      end
    else
      previous = item
    end
  end
  result << previous
  result
end

intervals = [[1,3],[6,9]]
new_interval = [2,5]
p insert(intervals, new_interval)

intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]]
new_interval = [4,8]
p insert(intervals, new_interval)

intervals = []
new_interval = [5,7]
p insert(intervals, new_interval)

