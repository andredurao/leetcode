# https://leetcode.com/problems/remove-vowels-from-a-string/

# @param {String} s
# @return {String}
def remove_vowels(s)
  s.gsub(/[aeiou]/, '')
end

p remove_vowels('leetcodeisacommunityforcoders')
p remove_vowels('aeiou')
