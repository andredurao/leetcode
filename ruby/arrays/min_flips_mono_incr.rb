# https://leetcode.com/problems/flip-string-to-monotone-increasing/

# @param {String} s
# @return {Integer}
def min_flips_mono_incr(s)
  m = 0
  s.each{|c| m += 1 if c == '0'}

  result = m

  s.each do |c|
    if c == '0'
      m -= 1
      result = [result, m].min
    else
      m += 1
    end
  end
  result
end
