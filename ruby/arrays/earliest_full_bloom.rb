# https://leetcode.com/problems/earliest-possible-day-of-full-bloom/
# @param {Integer[]} plant_time
# @param {Integer[]} grow_time
# @return {Integer}
def earliest_full_bloom(plant_time, grow_time)
  desc_indexes = []
  grow_time.each_with_index{|val,i| desc_indexes << [val,i]}
  desc_indexes.sort!{|x,y| y[0] <=> x[0]}
  # rebuild the original arrays both sorted desc by grow time
  new_plant_time = []
  new_grow_time = []
  desc_indexes.each do |_, i|
    new_plant_time << plant_time[i]
    new_grow_time << grow_time[i]
  end
  # prepending a 0 on grow_times
  # appending a 0 on plant_times
  new_plant_time << 0
  new_grow_time.unshift(0)

  plant_total = remainder = 0
  new_plant_time.zip(new_grow_time).each do |plant, grow|
    plant_total += plant
    if grow > 0
      remainder += grow - (plant + remainder)
      remainder = 0 if remainder < 0
    end
    # p [plant, grow, plant_total, remainder]
  end
  plant_total + remainder
end

# rest_day = 0
# more_necessary_days = []
# plant_time.zip(grow_time).sort_by(&:last).each do |plant_day, grow_day|
#   if rest_day < grow_day
#     more_necessary_days << (grow_day - rest_day)
#   end
#   rest_day += plant_day
# end
# plant_time.sum + more_necessary_days.max

plant_time = [1,4,3]
grow_time = [2,3,1]
p earliest_full_bloom(plant_time, grow_time)
p '----------------------------------------'
plant_time = [1,2,3,2]
grow_time = [2,1,2,1]
p earliest_full_bloom(plant_time, grow_time)
p '----------------------------------------'
plant_time = [1]
grow_time = [1]
p earliest_full_bloom(plant_time, grow_time)
p '----------------------------------------'
plant_time = [27,5,24,17,27,4,23,16,6,26,13,17,21,3,9,10,28,26,4,10,28,2]
grow_time = [26,9,14,17,6,14,23,24,11,6,27,14,13,1,15,5,12,15,23,27,28,12]
p earliest_full_bloom(plant_time, grow_time)
p '----------------------------------------'
