# https://leetcode.com/problems/guess-number-higher-or-lower/

# The guess API is already defined for you.
# @param num, your guess
# @return -1 if num is higher than the picked number
#          1 if num is lower than the picked number
#          otherwise return 0
# def guess(num)

def guessNumber(n)
  l = 1
  r = n

  while l <= r
    val = (l + r) / 2
    p [val, l, r]
    sleep 1
    compare = guess(val)
    return val if compare == 0
    if compare > 0
      l = val + 1
    else
      r = val - 1
    end
  end
  val
end

N = 6
def guess(val)
  if val > N
    -1
  elsif val < N
    1
  else
    0
  end
end

p guessNumber(10)
