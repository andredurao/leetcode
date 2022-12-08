# https://leetcode.com/problems/leaf-similar-trees/

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
# @param {TreeNode} root1
# @param {TreeNode} root2
# @return {Boolean}
def leaf_similar(root1, root2)
  vals1 = []
  leaves(root1, vals1)
  vals2 = []
  leaves(root2, vals2)
  vals1 == vals2
end

def leaves(root, vals)
  return if !root
  leaves(root.left, vals)
  vals << root.val if !root.left && !root.right
  leaves(root.right, vals)
end


tree1 = [3,5,1,6,2,9,8,nil,nil,7,4]
tree2 = [3,5,1,6,7,4,2,nil,nil,nil,nil,nil,nil,9,8]
root1 = Util.build(tree1)
root2 = Util.build(tree2)
p leaf_similar(root1, root2)

tree1 = [1,2,3]
tree2 = [1,3,2]
root1 = Util.build(tree1)
root2 = Util.build(tree2)
p leaf_similar(root1, root2)
