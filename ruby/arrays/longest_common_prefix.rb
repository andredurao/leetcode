# @param {String[]} strs
# @return {String}
def longest_common_prefix(strs)
  index = 0
  prefix = ''
  min_size = strs.map(&:size).min
  while index < min_size
    char = strs[0][index]

    same_char_at_index = strs.all? do |str|
      str[index] == char
    end

    if same_char_at_index
      prefix << char
      index += 1
    else
      return prefix
    end
  end
  prefix
end

strs = ["flower","flow","flight"]
p longest_common_prefix(strs)
strs = ["dog","racecar","car"]
p longest_common_prefix(strs)
strs = [""]
p longest_common_prefix(strs)
