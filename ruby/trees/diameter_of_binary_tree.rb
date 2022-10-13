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

def height(node)
  return 0 unless node
  l = height(node.left)
  r = height(node.right)
  l >= r ? (l+1) : (r+1)
end

def legspan(node)
  height(node.left) + height(node.right)
end

# @param {TreeNode} root
# @return {Integer}
def diameter_of_binary_tree(root)
  return 0 unless root
  l = diameter_of_binary_tree(root.left)
  r = diameter_of_binary_tree(root.right)
  [l, r, legspan(root)].max
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
p diameter_of_binary_tree(root)
# ===============================================
#       1
#     /
#   2
root = TreeNode.new(1)
root.left = TreeNode.new(2)
p diameter_of_binary_tree(root)
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
p diameter_of_binary_tree(root)
# ===============================================
#    3
#   /
#  1
#   \
#    2
root = TreeNode.new(3)
root.left = TreeNode.new(1)
root.left.right = TreeNode.new(2)
p diameter_of_binary_tree(root)
