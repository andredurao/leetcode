# https://leetcode.com/problems/integer-to-roman/description/

# @param {Integer} num
# @return {String}
def int_to_roman(num)
  algarisms = {
    'M' => 1000,
    'CM' => 900,
    'D' => 500,
    'CD' => 400,
    'C' => 100,
    'XC' => 90,
    'L' => 50,
    'XL' => 40,
    'X' => 10,
    'IX' => 9,
    'V' => 5,
    'IV' => 4,
    'I' => 1,
  }

  result = ''

  while num > 0 do
    algarisms.each do |symbol, value|
      if num >= value
        result += symbol
        num -= value
        break
      end
    end
  end

  result
end

p int_to_roman(3)
p int_to_roman(58)
p int_to_roman(1994)
p int_to_roman(20)
