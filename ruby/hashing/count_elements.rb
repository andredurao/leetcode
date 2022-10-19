require 'set'

# @param {Integer[]} arr
# @return {Integer}
def count_elements(arr)
  set = Set.new(arr)

  total = 0

  arr.each do |v|
    total += 1 if set.include?(v+1)
  end

  total
end

arr = [1,2,3]
p count_elements(arr)

arr = [1,1,3,3,5,5,7,7]
p count_elements(arr)

arr = [1,1,2,2]
p count_elements(arr)
