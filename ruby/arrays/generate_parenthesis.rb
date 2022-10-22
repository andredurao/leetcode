# @param {Integer} n
# @return {String[]}
def generate_parenthesis(n)
   results = []
   generate_parenthesis_helper(n, n, results)
   results
end

def generate_parenthesis_helper(l, r, results, str='')
  return if l < 0 || r < l

  if l == 0 && r == 0
    results << str
  else
    generate_parenthesis_helper(l-1, r, results, str + '(')
    generate_parenthesis_helper(l, r-1, results, str + ')')
  end
end

# Cracking the coding interview problem 8.9
p generate_parenthesis(1)
p generate_parenthesis(2)
p generate_parenthesis(3)
