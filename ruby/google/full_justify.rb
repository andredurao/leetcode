# @param {String[]} words
# @param {Integer} max_width
# @return {String[]}
def full_justify(words, max_width)
  lines = []
  while words.any?
    lines << fetch_line!(words, max_width)
  end

  result = []
  lines.each_with_index do |fetch_result, index|
    last = index == lines.size - 1
    line, length = fetch_result
    result << justify(line, max_width, length, last)
  end
  result
end

def fetch_line!(words, max_width)
  length = words.first.length
  line = [words.shift]
  while words.any? && (length + words.first.length < max_width)
    line << ' '
    word = words.shift
    line << word
    length += word.length + 1
  end
  [line, length]
end

def justify(line, max_width, length, last)
  if line.size == 1 || last
    spaces = max_width - line.join.length
    line << ' ' * spaces
    return line.join
  end
  # Q: how many spaces? A: the amount of odd value indexes
  # Lines will always have odd lengths, and size/2 spaces within
  spaces = line.size / 2

  even_division = (max_width - length) / spaces
  if even_division > 0
    index = -1
    line.map! do |word|
      index += 1
      append = index.even? ? '' : (' ' * even_division)
      word << append
    end
  end

  remaining = (max_width - length) % spaces
  n = 0
  while remaining > 0 # only fill left remaining spaces
    pos = 2 * n + 1
    line[pos] = line[pos] + ' '
    n += 1
    remaining -= 1
  end

  line.join
end

# max_width = 16
# words = ["This", "is", "an", "example", "of", "text", "justification."]
# result = full_justify(words, max_width)
# result.each{|x| p(x)}

# max_width = 16
# words = ["What","must","be","acknowledgment","shall","be"]
# result = full_justify(words, max_width)
# result.each{|x| p(x)}

max_width = 20
words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"]
result = full_justify(words, max_width)
result.each{|x| p(x)}

# line, len = fetch_line!(words, max_width)
# r = justify(line, max_width, len)
# p r
# binding.irb
