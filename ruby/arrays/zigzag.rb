# https://leetcode.com/problems/zigzag-conversion/

# @param {String} s
# @param {Integer} num_rows
# @return {String}
# from official solution
def convert(s, num_rows)
  return s if num_rows == 1
  result = ''
  group_size = 2 * (num_rows - 1)
  spaces = -1
  0.upto(num_rows - 1) do |row|
    i = row
    while i < s.size
      result << s[i]
      if ![0, num_rows - 1].include?(row)
        spaces = group_size - 2 * row
        j = i + spaces
        if j < s.size
          result << s[j]
        end
      end
      i += group_size
    end
  end
  result
end

s = "PAYPALISHIRING"
num_rows = 3
result = convert(s, num_rows)
p result

s = "PAYPALISHIRING"
num_rows = 4
result = convert(s, num_rows)
p result

s = "A"
num_rows = 1
result = convert(s, num_rows)
p result

s = "cbxdwjccgtdoqiscyspqzvuqivzptlpvooynyapgvswoaosaghrffnxnjyeeltzaizniccozwknwyhzgpqlwfkjqipuujvwtxlbznryjdohbvghmyuiggtyqjtmuqinntqmihntkddnalwnmsxsatqqeldacnnpjfermrnyuqnwbjjpdjhdeavknykpoxhxclqqedqavdwzoiorrwwxyrhlsrdgqkduvtmzzczufvtvfioygkvedervvudneghbctcbxdxezrzgbpfhzanffeccbgqfmzjqtlrsppxqiywjobspefujlxnmddurddiyobqfspvcoulcvdrzkmkwlyiqdchghrgytzdnobqcvdeqjystmepxcaniewqmoxkjwpymqorluxedvywhcoghotpusfgiestckrpaigocfufbubiyrrffmwaeeimidfnnzcphkflpbqsvtdwludsgaungfzoihbxifoprwcjzsdxngtacw"
num_rows = 472
result = convert(s, num_rows)
p result
