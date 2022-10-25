# @param {Integer[]} stones
# @return {Integer}
def last_stone_weight(stones)
  heap = []
  stones.each{|stone| heap_insert(heap, stone)}

  while heap.size > 1
    x = heap.size - 2
    y = heap.size - 1
    x_val = heap[x]
    y_val = heap[y]
    heap.delete_at(y)
    heap.delete_at(x)
    heap_insert(heap, y_val - x_val) if x_val != y_val
  end

  heap&.first || 0
end

def heap_insert(array, val)
  index = array.bsearch_index{|i| val < i} || array.size
  array.insert(index, val)
end

stones = [2,7,4,1,8,1]
p last_stone_weight(stones)
