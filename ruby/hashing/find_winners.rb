# @param {Integer[][]} matches
# @return {Integer[][]}
def find_winners(matches)
  winners = {}
  loosers = {}
  matches.each do |winner, looser|
    winners[winner] = true

    loosers[looser] ||= 0
    loosers[looser] += 1
  end

  loosers.keys.each do |looser|
    winners.delete(looser)
  end
  one_time_loosers = []
  loosers.each do |looser, qty|
     one_time_loosers << looser if qty == 1
  end
  [
    winners.keys.sort,
    one_time_loosers.sort
  ]
end
