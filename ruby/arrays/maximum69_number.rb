# @param {Integer} num
# @return {Integer}
def maximum69_number (num)
  digits = num.to_s.split('')
  first_6 = digits.index('6')
  if first_6
    digits[first_6] = '9'
    digits.join.to_i
  else
    num
  end
end

num = 9669
p maximum69_number(num)
