# @param {String} digits
# @return {String[]}
def letter_combinations(digits)
  result = []
  traverse(result, digits) if !digits.empty?
  result
end

def traverse(result, digits, string="")
  char_map = {
    '2' => %w(a b c),
    '3' => %w(d e f),
    '4' => %w(g h i),
    '5' => %w(j k l),
    '6' => %w(m n o),
    '7' => %w(p q r s),
    '8' => %w(t u v),
    '9' => %w(w x y z),
  }

  if digits.empty?
    result << string
    return
  end

  char_map[digits[0]].each do |char|
    traverse(result, digits[1..], string + char)
  end
end


digits = "23"
p letter_combinations(digits)


digits = ""
p letter_combinations(digits)
