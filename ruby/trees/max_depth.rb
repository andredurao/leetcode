# Definition for a binary tree node.
class TreeNode
  attr_accessor :val, :left, :right
  def initialize(val = 0, left = nil, right = nil)
      @val = val
      @left = left
      @right = right
  end
  def inspect
    "<TreeNode #{@val}>"
  end
end

# @param {TreeNode} root
# @return {Integer}
def max_depth(root)
  return -1 unless root
  l = max_depth(root.left)
  r = max_depth(root.right)
  [l+1, r+1].max
end

# ===============================================
#       1
#     /    \
#   2       3
#  / \
# 4   5
root = TreeNode.new(1)
root.left = TreeNode.new(2)
root.right = TreeNode.new(3)
root.left.left = TreeNode.new(4)
root.left.right = TreeNode.new(5)
# p max(root)
p max_depth(root)
# ===============================================
#       1
#     /
#   2
root = TreeNode.new(1)
root.left = TreeNode.new(2)
p max_depth(root)
# ===============================================
#         2
#        /
#       5
#      /
#     3
#    / \
#   1   4
root = TreeNode.new(2)
root.left = TreeNode.new(5)
root.left.left = TreeNode.new(3)
root.left.left.left = TreeNode.new(1)
root.left.left.right = TreeNode.new(4)
p max_depth(root)
# ===============================================
#    3
#   /
#  1
#   \
#    2
root = TreeNode.new(3)
root.left = TreeNode.new(1)
root.left.right = TreeNode.new(2)
p max_depth(root)
