# @param {String} s
# @return {Integer}
def min_deletions(s)
  counts = s.chars.tally.values.sort

  return 0 if [0,1].include?(counts.size)

  cursor = counts.size - 2
  total = 0

  # cursor goes starts on n-1 pos of the counts
  # and goes until 0 popping out repeated values
  while cursor >= 0
    p cursor, total, counts
    while counts[cursor] > 0 && counts[cursor] >= counts[cursor + 1]
      total += 1
      counts[cursor] -= 1
    end
    cursor -= 1
    p '----------------'
  end
  p cursor, total, counts


  total
end

def reduce_cursor?(counts, cursor)
  counts[cursor] != 0 && counts[cursor] == counts[cursor + 1]
end


s = "aaabbbcc"
# s = "aaabbbcccddd"
p  min_deletions(s)
