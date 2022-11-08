# https://leetcode.com/explore/interview/card/leetcodes-interview-crash-course-data-structures-and-algorithms/704/linked-lists/4598/
# Definition for singly-linked list.
# class ListNode
#     attr_accessor :val, :next
#     def initialize(val = 0, _next = nil)
#         @val = val
#         @next = _next
#     end
# end

require './list_node.rb'
require './util.rb'

# @param {ListNode} head
# @param {Integer} left
# @param {Integer} right
# @return {ListNode}
def reverse_between(head, left, right)
  stack = []
  rstack = []
  count = 1
  while head
    if count >= left && count <= right
      rstack.push(head)
      stack.push('X')
    else
      stack.push(head)
    end
    head = head.next
    count += 1
  end

  new_head = if stack[0] == 'X'
    rstack.pop
  else
    stack.shift
  end
  items = new_head
  while stack.any? || rstack.any?
    move = true
    if stack[0] == 'X'
      if rstack.empty?
        move = false
        stack.shift
      else
        items.next = rstack.pop
      end
    else
      items.next = stack.shift
    end
    items = items.next if move
  end
  items&.next = nil
  new_head
end

left = 2
right = 4
head = Util.deserialize([1,2,3,4,5])
result = reverse_between(head, left, right)
p Util.serialize(result)

left = 1
right = 1
head = Util.deserialize([5])
result = reverse_between(head, left, right)
p Util.serialize(result)

left = 1
right = 2
head = Util.deserialize([3,5])
result = reverse_between(head, left, right)
p Util.serialize(result)
