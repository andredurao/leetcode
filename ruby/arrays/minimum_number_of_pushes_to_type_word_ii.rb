# https://leetcode.com/problems/minimum-number-of-pushes-to-type-word-ii


# @param {String} word
# @return {Integer}
def minimum_pushes(word)
  # get the frequency of keys and sort it descending
  f = word.chars.tally
  chars = f.to_a.sort{|a,b| b[1] <=> a[1]}

  # build the keyboard combinations
  keymap = []
  8.times{ keymap << [] }
  slot = 0
  chars.each do |char, _|
    keymap[slot] << char
    slot += 1
    slot %= 8
  end

  # sum the total of keystrokes
  res = 0
  keymap.each do |key|
    key.each_with_index do |char, count|
      res += f[char] * (count+1)
    end
  end

  res
end

# -----------------------------


# word = "abcde"
# word = "xyzxyzxyzxyz"
word = "aabbccddeeffgghhiiiiii"
res = minimum_pushes(word)
puts(res)
