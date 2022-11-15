# https://leetcode.com/problems/count-complete-tree-nodes/description/
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
def count_nodes(root)
  # This is not performing as expected because it's running in O(n)
  return 0 unless root
  left = count_nodes(root.left)
  right = count_nodes(root.right)
  left + right + 1
end

tree = [1,2,3,4,5,6]
root = Util.build(tree)
p count_nodes(root)
