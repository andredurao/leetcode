# https://leetcode.com/problems/earliest-possible-day-of-full-bloom/
# @param {Integer[]} plant_time
# @param {Integer[]} grow_time
# @return {Integer}
def earliest_full_bloom(plant_time, grow_time)
  grow_time_by_indexes = []
  grow_time.each_with_index{|val, index| grow_time_by_indexes << [val, index]}
  grow_time_by_indexes.sort!{|x,y| y[0] <=> x[0]}
  total_plant_time = total = diff = 0
  grow_time_by_indexes.each_with_index do |indexes, current_index|
    index, meta_index = indexes
    total_plant_time += plant_time[meta_index]
    total = [total, total_plant_time + grow_time[meta_index]].max
  end
  total
end


plant_time = [1,4,3]
grow_time = [2,3,1]
p earliest_full_bloom(plant_time, grow_time)


plant_time = [1,2,3,2]
grow_time = [2,1,2,1]
p earliest_full_bloom(plant_time, grow_time)


plant_time = [1]
grow_time = [1]
p earliest_full_bloom(plant_time, grow_time)


plant_time = [27,5,24,17,27,4,23,16,6,26,13,17,21,3,9,10,28,26,4,10,28,2]
grow_time = [26,9,14,17,6,14,23,24,11,6,27,14,13,1,15,5,12,15,23,27,28,12]
p earliest_full_bloom(plant_time, grow_time)
