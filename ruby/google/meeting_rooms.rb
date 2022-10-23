require 'set'

# @param {Integer[][]} intervals
# @return {Integer}
def min_meeting_rooms(intervals)
  # sort intervals by their start times
  intervals.sort!{|x,y| x <=> y}
  set = Set.new()
  start_indexes = {}
  end_indexes = {}

  intervals.each_with_index do |interval, index|
    start_time = interval[0]
    end_time = interval[1]

    set.add(start_time)
    set.add(end_time)

    # map the qty of start and ends at each value
    start_indexes[start_time] ||= 0
    end_indexes[end_time] ||= 0
    start_indexes[start_time] = start_indexes[start_time] + 1
    end_indexes[end_time] = end_indexes[end_time] + 1
  end


  # search for max use of rooms at the same time
  values = set.to_a.sort
  in_use = max_in_use = 0
  values.each do |value|
    p [value, in_use]
    in_use += (start_indexes[value] || 0)
    in_use -= (end_indexes[value] || 0)
    max_in_use = [in_use, max_in_use].max
  end

  max_in_use
end

intervals = [[0,30],[5,10],[15,20]]
p min_meeting_rooms(intervals)

intervals = [[7,10],[2,4]]
p min_meeting_rooms(intervals)
