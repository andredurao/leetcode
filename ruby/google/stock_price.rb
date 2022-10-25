class StockPrice
  def initialize()
    @timestamps_map = {}
    @prices_map = {}
    @current_timestamp = 0
  end

=begin
  :type timestamp: Integer
  :type price: Integer
  :rtype: Void
=end
  def update(timestamp, price)
    if @timestamps_map[timestamp]
      old_price = @timestamps_map[timestamp]
      count = @prices_map[old_price]
      if count == 1
        @prices_map.delete(old_price)
      else
        @prices_map[old_price] = count - 1
      end
    end
    @timestamps_map[timestamp] = price
    @prices_map[price] ||= 0
    @prices_map[price] = @prices_map[price] + 1
    @prices_map = @prices_map.sort.to_h
    @current_timestamp = [@current_timestamp, timestamp].max
  end

=begin
  :rtype: Integer
=end
  def current()
    @timestamps_map[@current_timestamp]
  end

=begin
  :rtype: Integer
=end
  def maximum()
    @prices_map.keys.last
  end

=begin
:rtype: Integer
=end
  def minimum()
    @prices_map.keys.first
  end

end

# Your StockPrice object will be instantiated and called as such:
# obj = StockPrice.new()
# obj.update(timestamp, price)
# param_2 = obj.current()
# param_3 = obj.maximum()
# param_4 = obj.minimum()

input1 = ["StockPrice","update","update","current","maximum","update","maximum","update","minimum"]
input2 = [[],[1,10],[2,5],[],[],[1,3],[],[4,2],[]]
s = StockPrice.new
s.update(1,10)
s.update(2,5)
p s.current
p s.maximum
s.update(1,3)
p s.maximum
s.update(4,2)
p s.minimum
