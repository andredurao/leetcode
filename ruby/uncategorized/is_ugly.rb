# https://leetcode.com/problems/ugly-number/

# @param {Integer} n
# @return {Boolean}
def is_ugly(n)
  factors = [2,3,5]
  while n > 0
    return true if n == 1
    count = 0
    factors.each do |factor|
      if n % factor == 0
        n /= factor
        count += 1
      end
    end
    return false if count == 0
  end
  false
end

p is_ugly(6)
p is_ugly(1)
p is_ugly(14)

