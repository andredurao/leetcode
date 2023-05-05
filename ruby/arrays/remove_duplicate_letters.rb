# https://leetcode.com/problems/remove-duplicate-letters/

require 'set'

# @param {String} s
# @return {String}
# @param {String} s
# @return {String}
def remove_duplicate_letters(s)
  seen = Set.new
  q = []
  result = ''

  frequencies = s.chars.tally

  s.chars.each_with_index do |c, i|
    frequencies[c] -= 1
    if !seen.include?(c)
      while q.any? && c < q.last && frequencies[q.last] > 0
        seen.delete(q.pop)
      end
      seen.add(c)
      q << c
    end
  end

  q.join
end

s = "bcabc" # abc
p remove_duplicate_letters(s)

s = "cbacdcbc" # acbd
p remove_duplicate_letters(s)
