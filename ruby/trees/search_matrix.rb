# @param {Integer[][]} matrix
# @param {Integer} target
# @return {Boolean}
def search_matrix(matrix, target)
  height = matrix.size
  width = matrix[0].size
  l = 0
  r = height * width - 1
  while l <= r
    mid = (l+r) / 2
    current = val(matrix, mid, width)
    # p [l,r,mid,current]
    return true if current == target
    if current < target
      l = mid + 1
    else current > target
      r = mid - 1
    end
  end
  false
end

def val(matrix, index, width)
  matrix[index / width][index % width]
end

matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]]
target = 3
p search_matrix(matrix, target)

matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]]
target = 13
p search_matrix(matrix, target)

matrix = [[1,1]]
target = 2
p search_matrix(matrix, target)

matrix = [[1]]
target = 1
p search_matrix(matrix, target)

matrix = [[1,1]]
target = 0
p search_matrix(matrix, target)

matrix = [[1],[3]]
target = 3
p search_matrix(matrix, target)
