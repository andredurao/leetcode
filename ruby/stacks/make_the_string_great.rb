# @param {String} s
# @return {String}
def make_good(s)
  queue = []
  s.chars.each do |char|
    last = queue.last
    if last != char && ((char.upcase == last) || (char.downcase == last))
      queue.pop
    else
      queue.push(char)
    end
  end
  queue.join
end

s = "leEeetcode"
p make_good(s)

s = "abBAcC"
p make_good(s)
