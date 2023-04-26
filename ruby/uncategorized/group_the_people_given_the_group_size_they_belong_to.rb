# https://leetcode.com/problems/group-the-people-given-the-group-size-they-belong-to/description/

# @param {Integer[]} group_sizes
# @return {Integer[][]}
def group_the_people(group_sizes)
  map = {}

  group_sizes.each_with_index do |size, index|
    map[size] ||= [[]]
    map[size] << [] if map[size][-1].size == size
    map[size][-1] << index
  end

  sizes = map.keys.sort
  result = []
  sizes.each{|size| result += map[size]}

  result
end
