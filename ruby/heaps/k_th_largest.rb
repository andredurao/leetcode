class KthLargest

  attr_accessor :nums
  def initialize(k, nums)
    @k = k
    @nums = []
    nums.each{|val| add(val)}
  end

=begin
  :type val: Integer
  :rtype: Integer
=end
  def add(val)
    index = @nums.bsearch_index{|i| val < i} || @nums.size
    @nums.insert(index, val)
    @nums[(@k) * -1]
  end
end

kthLargest = KthLargest.new(3, [4, 5, 8, 2]);
# p kthLargest.nums
p kthLargest.add(3)
p kthLargest.add(5)
p kthLargest.add(10)
p kthLargest.add(9)
p kthLargest.add(4)
