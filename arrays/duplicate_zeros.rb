# @param {Integer[]} arr
# @return {Void} Do not return anything, modify arr in-place instead.
def duplicate_zeros(arr)
  copy = []
  index = 0
  while index < arr.length && copy.length < arr.length
    copy << arr[index]
    if arr[index] == 0 && copy.length < arr.length
      copy << 0
    end
    index += 1
  end
  index = -1
  arr.map! do |value|
    index += 1
    copy[index]
  end
end

arr = [1,0,2,3,0,4,5,0]

duplicate_zeros(arr)
p arr
