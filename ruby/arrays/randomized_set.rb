# https://leetcode.com/problems/insert-delete-getrandom-o1/

class RandomizedSet
  attr_accessor :values, :list

  def initialize()
    @values = {}
    @list = []
  end

=begin
  :type val: Integer
  :rtype: Boolean
=end
  def insert(val)
    # p ['+', val, values, list]
    return false if values[val]

    values[val] = list.size
    list << val
    true
  end

=begin
  :type val: Integer
  :rtype: Boolean
=end
  def remove(val)
    # p ['-', val, values, list]
    if index = values[val]
      last = list.last
      list[index] = last
      values[last] = index
      list.pop
      values.delete(val)
      true
    else
      false
    end
  end

=begin
  :rtype: Integer
=end
  def get_random()
    list.sample
  end
end
