# https://leetcode.com/problems/binary-tree-preorder-traversal/description/
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
# @return {Integer[]}
def preorder_traversal(root)
  pre_traversal_aux(root)
end

def pre_traversal_aux(node, items=[])
  return items unless node
  items << node.val
  pre_traversal_aux(node.left, items) if node.left
  pre_traversal_aux(node.right, items) if node.right
  return items
end

tree = [1,nil,2,3]
root = Util.build(tree)
p preorder_traversal(root)

tree = []
root = Util.build(tree)
p preorder_traversal(root)


tree = [1]
root = Util.build(tree)
p preorder_traversal(root)
