# https://leetcode.com/problems/find-the-difference/description/

# @param {String} s
# @param {String} t
# @return {Character}
def find_the_difference(s, t)
  map_s = s.chars.tally
  map_t = t.chars.tally
  map_t.each do |char, count|
    return char if map_s[char] != count
  end
end

s = "abcd"
t = "abcde"
p find_the_difference(s, t)
