# https://leetcode.com/problems/online-stock-span/
# maximum number of consecutive days (starting from today and going backward)
# for which the stock price was less than or equal to today's price

class StockSpanner
  def initialize()
    @prices = []
    @totals = {}
  end

=begin
    :type price: Integer
    :rtype: Integer
=end
  def next(price)
    current = -1
    total = 1
    while @prices.any? && @prices.last <= price
      current = @prices.pop
      total += @totals[current]
    end
    @prices << price
    @totals[price] = total
    total
  end
end

# stockSpanner = StockSpanner.new
# p stockSpanner.next(100) # return 1
# p stockSpanner.next(80)  # return 1
# p stockSpanner.next(60)  # return 1
# p stockSpanner.next(70)  # return 2
# p stockSpanner.next(60)  # return 1
# p stockSpanner.next(75)  # return 4, because the last 4 prices (including today's price of 75) were less than or equal to today's price.
# p stockSpanner.next(85)  # return 6

p '[null,1,1,2,3,5,1,1,3,1,5,6,1,1,3,15]'
stockSpanner = StockSpanner.new

p stockSpanner.next(90) # 90 1
p stockSpanner.next(21) # 21 2
p stockSpanner.next(21) # 68 3
p stockSpanner.next(68) # 94 5
p stockSpanner.next(94) # 13 1
p stockSpanner.next(13) # 1  1
p stockSpanner.next(1 ) # 37 3
p stockSpanner.next(37) # 3  1
p stockSpanner.next(3 ) # 61
p stockSpanner.next(61) # 86
p stockSpanner.next(86) # 19
p stockSpanner.next(19) # 12
p stockSpanner.next(12) # 35
p stockSpanner.next(35) # 96
p stockSpanner.next(96) #
