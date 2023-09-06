# https://leetcode.com/problems/split-linked-list-in-parts/?envType=daily-question&envId=2023-09-06

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
# @param {ListNode} head
# @param {Integer} k
# @return {ListNode[]}
def split_list_to_parts(head, k)
  size = list_size(head)
  packet_size = size / k
  extras = size % k
  result = []

  k.times do
    head ||= ListNode.new(nil)
    start = head
    binding.irb
    packet = []
    qty = packet_size
    if extras > 0
      qty += 1
      extras -= 1
    end
    (qty - 1).times do
      head = head&.next || ListNode.new(nil)
    end
    next_node = head&.next || ListNode.new(nil)
    head.next = nil
    result << start
    head = next_node
  end
  result.each do |x|
    p "---"
    p x
  end
  result
end

def list_size(head)
  result = 0
  while head
    result += 1
    head = head.next
  end
  result
end
# ---------------------------------

head = Util.deserialize([])
k = 3

p split_list_to_parts(head, k)
