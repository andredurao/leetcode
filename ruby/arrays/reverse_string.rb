# @param {Character[]} s
# @return {Void} Do not return anything, modify s in-place instead.
def reverse_string(s)
  s.reverse!
end

s = %w[h e l l o]
r = reverse_string(s)
p r
