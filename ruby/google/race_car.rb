# @param {Integer} target
# @return {Integer}
def racecar(target)
  limit = target.bit_length + 1
  barrier = 1 << limit
  pq = [[0, target]]
  dist = [Float::INFINITY] * (2 * barrier + 1)
  dist[target] = 0
  # binding.irb

  while pq.any?
    steps, targ = pq.pop
    next if targ && dist[targ] && dist[targ] > steps

    0.upto(limit+1) do |k|
      walk = (1 << k) - 1
      steps2 = steps + k + 1
      targ2 = walk - targ
      steps2 -= 1 if walk == targ  # No "R" command if already exact

      if targ2.abs <= barrier && steps2 < dist[targ2]
        heap_push(pq, [steps2, targ2])
        dist[targ2] = steps2
      end
    end
  end

  return dist[0]
end

def heap_push(arr, x)
  index = arr.bsearch_index { |y| (y[0] <=> x[0]) >= 0 } || arr.size
  arr.insert(index, x)
end

target = 330
p racecar(target)

# # https://leetcode.com/problems/minimize-deviation-in-array/discuss/1053053/Ruby-Working-with-Priority-Queue-in-Ruby-one-approach/842794/
# # Insert element into ASC sorted array.
# def heappush(arr, x)
#   index = arr.bsearch_index { |y| (y <=> x) >= 0 } || arr.size
#   arr.insert(index, x)
# end
# # Remove element from ASC sorted array if the element is present.
# def heappop(arr, x)
#   index = arr.bsearch_index { |y| (y <=> x) >= 0 }
#   arr.delete_at(index) if index && arr[index] == x
# end
