# @param {Integer[]} people
# @param {Integer} limit
# @return {Integer}
def num_rescue_boats(people, limit)
  freq = people.tally
  weights = freq.keys.sort
  l = 0
  r = weights.size - 1
  total = 0

  while l <= r
    pack = []
    sum = 0
    while sum <= limit && l <= r && pack.size < 2
      used = 0

      # consume right
      if freq[weights[r]] > 0 && sum + weights[r] <= limit
        pack << weights[r]
        sum += weights[r]
        freq[weights[r]] -= 1
        used += 1
      end

      # consume left
      if freq[weights[l]] > 0 && sum + weights[l] <= limit && used == 0
        pack << weights[l]
        sum += weights[l]
        freq[weights[l]] -= 1
        used += 1
      end

      # move left pointer
      l += 1 if freq[weights[l]] == 0
      # move right pointer
      r -= 1 if freq[weights[r]] == 0

      break if used == 0
    end
    # puts "pack = #{pack} l [#{l}] r [#{r}] freq [#{freq}]"
    total += 1
  end

  total
end

# people = [11,2,8,1]
# limit = 11
people = [11,2,2,8,8]
limit = 11
# people = [3,2,3,2,2]
# limit = 6

puts num_rescue_boats(people, limit)
