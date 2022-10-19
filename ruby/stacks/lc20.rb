# @param {String} s
# @return {Boolean}
def is_valid(s)
  map = { ')' => '(', ']' => '[', '}' => '{' }

  stack = []
  s.chars.each do |char|
    if map[char]
      return false if stack.pop != map[char]
    else
      stack.push(char)
    end
  end

  stack.empty?
end


s = "()"
p is_valid(s)

s = "()[]{}"
p is_valid(s)

s = "(]"
p is_valid(s)
