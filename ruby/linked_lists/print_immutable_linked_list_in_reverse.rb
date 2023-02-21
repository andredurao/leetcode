# https://leetcode.com/problems/print-immutable-linked-list-in-reverse/
require './list_node.rb'
require './util.rb'

#   This is the ImmutableListNode's API interface.
#   You should not implement it, or speculate about its implementation.
#
#   class ImmutableListNode
#      def printValue()
# .        print the value of this node.
#      def end
#          """
#
#      def getNext()
# .        return the next node.
#      end
#   end

def printLinkedListInReverse(head)
  return unless head
  printLinkedListInReverse(head.getNext())
  head.printValue()
end

head = Util.deserialize([1,2,3,4])
p printLinkedListInReverse(head)
