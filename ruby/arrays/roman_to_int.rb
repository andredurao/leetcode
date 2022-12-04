# https://leetcode.com/problems/roman-to-integer/

# @param {String} s
# @return {Integer}
def roman_to_int(s)
  map = {
    'M'  => 1000,
    'CM' => 900,
    'D'  => 500,
    'CD' => 400,
    'C'  => 100,
    'XC' => 90,
    'L'  => 50,
    'XL' => 40,
    'X'  => 10,
    'IX' => 9,
    'V'  => 5,
    'IV' => 4,
    'I'  => 1,
  }

  result = 0
  while s.length > 0
    map.each do |k,v|
      if s.index(k) == 0
        result += v
        s = s[k.length..-1]
        break
      end
    end
  end
  result
end

s = "III"
p roman_to_int(s)
s = "LVIII"
p roman_to_int(s)
s = "MCMXCIV"
p roman_to_int(s)
