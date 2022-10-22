require './list_node.rb'
require './util.rb'

# @param {ListNode} head
# @return {ListNode}
def reverse_list(head)
  return unless head
  stack = []
  while head
    stack.push head
    head = head.next
  end
  new_head = stack.pop
  items = new_head
  while stack.any?
    items.next = stack.pop
    items = items.next
  end
  items.next = nil
  new_head
end

# ---------------------------------



head = Util.deserialize([1,2,3,4,5])
reverse = reverse_list(head)
p Util.serialize(reverse)
