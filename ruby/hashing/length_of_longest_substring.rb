# @param {String} s
# @return {Integer}
def length_of_longest_substring(s)
  map = {}
  max = total = start = index = 0
  while index < s.length
    char = s[index]
    if map[char]
      start = [map[char], start].max
      map[char] = index
      total = index - start
    else
      total += 1
      map[char] = index
    end
    max = [total, max].max
    index += 1
  end
  max
end

strs = [
  "abcabcbb",
  "bbbbb",
  "pwwkew"
]
strs.each{|s| p length_of_longest_substring(s)}

