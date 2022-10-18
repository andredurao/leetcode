# @param {String} sentence
# @return {Boolean}
def check_if_pangram(sentence)
  map = {}
  %w[a b c d e f g h i j k l m n o p q r s t u v w x y z].each{|char| map[char] = 0}
  sentence.chars.each{|char| map[char] += 1}
  map.all?{|char, count| count > 0}
end

sentence = 'xxx'
p check_if_pangram(sentence)
sentence = 'abcdefghijklmnopqrstuvwxyz'
p check_if_pangram(sentence)
