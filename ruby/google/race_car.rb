# @param {Integer} target
# @return {Integer}
def racecar(target)
  queue = [{pos: 0, speed: 1, steps: 0}]
  visited = {}
  while queue.any?
    state = queue.pop
    return state[:steps] if state[:pos] == target

    if (state[:pos] + state[:speed]) <= 2 * target
      new_state = {
        pos: state[:pos] + state[:speed],
        speed: 2 * state[:speed],
        steps: state[:steps] + 1
      }
      key = "#{new_state[:pos]}-#{new_state[:speed]}"
      if !visited[key]
        visited[key] = true
        queue.push(new_state)
      end
    end
    if state[:pos] >= (target / 2)
      new_state = {
        pos: state[:pos],
        speed: (state[:speed] == 0 ? -1 : 1),
        steps: state[:steps] + 1
      }
      key = "#{new_state[:pos]}-#{new_state[:speed]}"
      if !visited[key]
        visited[key] = true
        # queue.push(new_state)
        heap_push(queue, new_state)
      end
    end
  end
  return -1
end

def heap_push(arr, x)
  index = arr.bsearch_index { |y| (y[:pos] <=> x[:pos]) >= 0 } || arr.size
  arr.insert(index, x)
end

target = ARGV[0].to_i
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
