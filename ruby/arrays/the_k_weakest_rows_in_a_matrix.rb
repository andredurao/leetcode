# https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/

# @param {Integer[][]} mat
# @param {Integer} k
# @return {Integer[]}
def k_weakest_rows(mat, k)
  list = []
  mat.each_index do |i|
    count = mat[i].tally[1] || 0
    list << [i, count]
  end

  list.sort! do |a,b|
    if a[1] != b[1]
      a[1] <=> b[1]
    else
      a[0] <=> b[0]
    end
  end
  list.map!{|x| x[0]}[..(k - 1)]
end

mat = [[1,0],[0,0],[1,0]]
k = 2
p k_weakest_rows(mat, k)
