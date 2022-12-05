# https://leetcode.com/problems/merge-two-sorted-lists
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
# @param {ListNode} list1
# @param {ListNode} list2
# @return {ListNode}
def merge_two_lists(list1, list2)
  result = nil
  while list1 || list2
    if !list2
      result = append(result, list1.val)
      list1 = list1.next
    elsif !list1
      result = append(result, list2.val)
      list2 = list2.next
    elsif list1.val <= list2.val
      result = append(result, list1.val)
      list1 = list1.next
    elsif list2.val < list1.val
      result = append(result, list2.val)
      list2 = list2.next
    end
  end
  result
end

def append(result, val)
  node = ListNode.new(val)
  el = result
  if !el
    result = node
  else
    while el.next
      el = el.next
    end
    el.next = node
  end
  result
end

# list1 = Util.deserialize([1,2,4])
# list2 = Util.deserialize([1,3,4])
# result = merge_two_lists(list1, list2)
# p Util.serialize(result)


# list1 = Util.deserialize([])
# list2 = Util.deserialize([])
# result = merge_two_lists(list1, list2)
# p Util.serialize(result)

list1 = Util.deserialize([])
list2 = Util.deserialize([0])
result = merge_two_lists(list1, list2)
p Util.serialize(result)
