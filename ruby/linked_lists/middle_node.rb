# Definition for singly-linked list.
class ListNode
  attr_accessor :val, :next
  def initialize(val = 0, _next = nil)
    @val = val
    @next = _next
  end
end

# @param {ListNode} head
# @return {ListNode}
def middle_node(head)
  length = 1
  item = head
  while item.next
    length += 1
    item = item.next
  end

  item = head
  n = 1
  while n < (length / 2) + 1
    n += 1
    item = item.next
  end
  item
end

def populate(list)
  head = ListNode.new(list.shift)
  _next = nil
  list.each do |val|
    _next = head unless _next
    _next.next = ListNode.new(val)
    _next = _next.next
  end

  head
end

head = populate([1,2,3,4,5])
p middle_node(head)
