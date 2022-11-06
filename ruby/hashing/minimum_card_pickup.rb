# @param {Integer[]} cards
# @return {Integer}
def minimum_card_pickup(cards)
  map = {}
  cards.each_with_index do |card, index|
    map[card] ||= []
    map[card] = (map[card] + [index]).sort{|x,y| y <=> x}
  end
  min = Float::INFINITY
  map.each do |card, indexes|
    next if indexes.size < 2
    # slide in pairs
    delta = 0
    while delta < indexes.size - 1
      min = [min, indexes[delta] - indexes[delta+1]].min
      delta += 1
    end
  end
  if min == Float::INFINITY
    -1
  else
    min + 1
  end
end

cards = [3,4,2,3,4,7]
p cards
p minimum_card_pickup(cards)

cards = [1,0,5,3]
p cards
p minimum_card_pickup(cards)
