# @param {Integer[]} arr
# @return {Void} Do not return anything, modify arr in-place instead.
# Complexity analysis:
# Runtime: O(n), Space: O(n)
def duplicate_zeros(arr)
  queue = []
  arr.map! do |val|
    # to duplicate the element
    queue.push(val) if val == 0

    if queue.empty?
      val
    else
      queue.push(val)
      queue.shift
    end
  end
end

arr = [1,0,2,3,0,4,5,0]

duplicate_zeros(arr)
p arr
