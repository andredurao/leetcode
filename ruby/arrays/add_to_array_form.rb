# @param {Integer[]} num
# @param {Integer} k
# @return {Integer[]}
def add_to_array_form(num, k)
  val = num.map(&:to_s).join.to_i + k
  val.to_s.chars.map(&:to_i)
end
