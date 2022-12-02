# https://leetcode.com/problems/determine-if-two-strings-are-close/

# @param {String} word1
# @param {String} word2
# @return {Boolean}
def close_strings(word1, word2)
  freq1 = word1.chars.sort.tally
  freq2 = word2.chars.sort.tally
  return false if freq1.keys != freq2.keys

  chars1 = Array.new(26, 0)
  chars2 = Array.new(26, 0)
  freq1.each do |key, value|
    chars1[key.ord - 97] = value
    chars2[key.ord - 97] = freq2[key]
  end
  chars1.sort == chars2.sort
end

word1 = "abc"
word2 = "bca"
p close_strings(word1, word2)

word1 = "cabbba"
word2 = "abbccc"
p close_strings(word1, word2)
