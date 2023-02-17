# https://leetcode.com/problems/minimum-distance-between-bst-nodes/
require './util.rb'
require './tree.rb'

# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val = 0, left = nil, right = nil)
#         @val = val
#         @left = left
#         @right = right
#     end
# end
# @param {TreeNode} root
# @return {Integer}
def min_diff_in_bst(root)
  items = in_order_traversal(root)
  min_diff = Float::INFINITY
  prev = nil
  items.each do |item|
    if prev
      min_diff = [(prev - item).abs, min_diff].min
      return min_diff if min_diff == 1
    end
    prev = item
  end
  min_diff
end

def in_order_traversal(root, items=[])
  in_order_traversal(root.left, items) if root.left
  items << root.val
  in_order_traversal(root.right, items) if root.right
  items
end

tree = [4,2,6,1,3]
root = Util.build(tree)
p min_diff_in_bst(root)

tree = [1,0,48,nil,nil,12,49]
root = Util.build(tree)
p min_diff_in_bst(root)
