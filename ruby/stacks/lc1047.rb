# @param {String} s
# @return {String}
def remove_duplicates(s)
  stack = []

  s.chars.each do |c|
    if c == stack.last
      stack.pop
    else
      stack.push(c)
    end
  end
  stack.join
end

s = "abbaca"
p remove_duplicates(s)
