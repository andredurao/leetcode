require './list_node.rb'
require './util.rb'


# Definition for singly-linked list.
# class ListNode
#     attr_accessor :val, :next
#     def initialize(val = 0, _next = nil)
#         @val = val
#         @next = _next
#     end
# end
# @param {ListNode} l1
# @param {ListNode} l2
# @return {ListNode}
def add_two_numbers(l1, l2)
  return ListNode.new(nil) if !(l1 && l2)

  getval = ->(node){ node&.val || 0 }

  result = nil ; head = nil
  acc = 0
  while l1&.val || l2&.val do
    total = getval.(l1) + getval.(l2) + acc
    acc = total / 10
    val = total % 10
    if result
      result.next = ListNode.new(val)
      result = result.next
    else
      result = ListNode.new(val)
      head = result
    end
    l1 = l1&.next || ListNode.new(nil)
    l2 = l2&.next || ListNode.new(nil)
  end
  result.next = ListNode.new(acc) if acc > 0


  head
end

# ---------------------------------
l1 = Util.deserialize([2,4,3])
l2 = Util.deserialize([5,6,4])
# p add_two_numbers(l1, l2)
# ---------------------------------
# p add_two_numbers(nil, nil)
# ---------------------------------
l1 = Util.deserialize([9,9,9,9,9,9,9])
l2 = Util.deserialize([9,9,9,9])
p add_two_numbers(l1, l2)
