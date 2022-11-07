# @param {String} jewels
# @param {String} stones
# @return {Integer}
def num_jewels_in_stones(jewels, stones)
  jewels_map = jewels.split('').tally
  total = 0
  stones.split('').each{|stone| jewels[stone] && total += 1}
  total
end
