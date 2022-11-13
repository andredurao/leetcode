# https://leetcode.com/problems/reverse-words-in-a-string/
# @param {String} s
# @return {String}
def reverse_words(s)
  s.split.reverse.join(' ')
end

s = "the sky is blue"
p reverse_words(s)
