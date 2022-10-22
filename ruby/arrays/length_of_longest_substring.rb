# @param {String} s
# @return {Integer}
# there's a ressemblence with boyer moore's algorithm
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


s = "abcabcbb"
p length_of_longest_substring(s)
s = "bbbbb"
p length_of_longest_substring(s)
s = "pwwkew"
p length_of_longest_substring(s)
s = "abba"
p length_of_longest_substring(s)
s = "aab"
p length_of_longest_substring(s)
s = "abba"
p length_of_longest_substring(s)

# # ----
s = "abcdefghijklmnopqrstuvwxyz"
p length_of_longest_substring(s)
s = "abcdefbghijklmnopqrstuvwxyz"
p length_of_longest_substring(s)
