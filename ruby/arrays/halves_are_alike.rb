# https://leetcode.com/problems/determine-if-string-halves-are-alike

# @param {String} s
# @return {Boolean}
def halves_are_alike(s)
  half = s.size / 2
  s1 = s[0...half].upcase.scan(/[AEIOU]/).size
  s2 = s[half..-1].upcase.scan(/[AEIOU]/).size
  s1 == s2
end

s = "book"
p halves_are_alike(s)

s = "textbook"
p halves_are_alike(s)
