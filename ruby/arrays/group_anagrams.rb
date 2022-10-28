# @param {String[]} strs
# @return {String[][]}
def group_anagrams(strs)
  anagrams_map = {}
  strs.each do |str|
    chars = str.chars.sort.join
    anagrams_map[chars] ||= []
    anagrams_map[chars] << str
  end
  anagrams_map.values
end

strs = ["eat","tea","tan","ate","nat","bat"]
p group_anagrams(strs)
