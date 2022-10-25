class MedianFinder
  def initialize()
    @values = []
  end


=begin
  :type num: Integer
  :rtype: Void
=end
  def add_num(num)
    index = @values.bsearch_index{|i| num < i } || @values.size
    @values.insert(index, num)
  end


=begin
  :rtype: Float
=end
  def find_median()
    center = @values.size / 2
    if @values.size.odd?
      @values[center]
    else
      @values[(center - 1)..center].sum / 2.0
    end
  end


end

# Your MedianFinder object will be instantiated and called as such:
mf = MedianFinder.new
mf.add_num(1)
mf.add_num(2)
p mf.find_median
