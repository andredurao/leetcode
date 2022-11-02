# @param {String} s
# @param {String} t
# @return {Boolean}
def is_subsequence(s, t)
  cursor = 0
  t.chars.each do |c|
    cursor += 1 if c == s[cursor]
    return true if cursor == s.size
  end
  cursor == s.size
end

s = "abc"
t = "ahbgdc"
p is_subsequence(s, t)
