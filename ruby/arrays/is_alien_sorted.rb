# https://leetcode.com/problems/verifying-an-alien-dictionary/

# @param {String[]} words
# @param {String} order
# @return {Boolean}
def is_alien_sorted(words, order)
  map = order.chars.map.with_index.to_h

  0.upto(words.size - 2) do |i|
    current_word = words[i]
    next_word = words[i + 1]
    current_word.chars.each_with_index do |current_char, j|
      next_char = next_word[j]
      # p [current_char, next_char, map[current_char], map[next_char]]
      return false if !next_char || map[current_char] > map[next_char]
      break if !current_char || map[current_char] < map[next_char]
    end
  end
  true
end

words = ["hello","leetcode"]
order = "hlabcdefgijkmnopqrstuvwxyz"
p is_alien_sorted(words, order)

words = ["word","world","row"]
order = "worldabcefghijkmnpqstuvxyz"
p is_alien_sorted(words, order)

words = ["apple","app"]
order = "abcdefghijklmnopqrstuvwxyz"
p is_alien_sorted(words, order)

words = ["kuvp","q"]
order = "ngxlkthsjuoqcpavbfdermiywz"
p is_alien_sorted(words, order)
