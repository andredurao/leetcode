require './list_node.rb'
require './util.rb'

# https://leetcode.com/problems/linked-list-cycle
# Definition for singly-linked list.
# class ListNode
#     attr_accessor :val, :next
#     def initialize(val)
#         @val = val
#         @next = nil
#     end
# end
# @param {ListNode} head
# @return {Boolean}
def hasCycle(head)
  while head
    # that may be a hack, for some, but the lang allows us to add new instance vars :shrug:
    return true if head.instance_variable_defined?(:@visited)
    head.instance_variable_set(:@visited, true)
    head = head.next
  end
  false
end
