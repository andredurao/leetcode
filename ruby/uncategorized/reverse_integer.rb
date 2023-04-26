# https://leetcode.com/problems/reverse-integer/description/

# @param {Integer} x
# @return {Integer}
def reverse(x)
  result = 0
  negative = false
  if x < 0
    negative = true
    x *= -1
  end

  while x > 0
    result *= 10
    result += x % 10
    x /= 10
  end

  result *= -1 if negative
  result = 0 if overflow?(result)

  result
end

def overflow?(x)
  lower_limit = -(2 << 30)
  higher_limit = (2 << 30) - 1
  x > higher_limit || x < lower_limit
end
