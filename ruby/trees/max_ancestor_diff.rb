# Definition for a binary tree node.
class TreeNode
  attr_accessor :val, :left, :right
  def initialize(val = 0, left = nil, right = nil)
    @val = val
    @left = left
    @right = right
  end
end

# @param {TreeNode} root
# @return {Integer}
def max_ancestor_diff(root)
  return 0 unless root
  max_ancestor_diff_aux(root, root.val, root.val)
end

def max_ancestor_diff_aux(node, max, min)
  p [node&.val, max, min]
  return max - min unless node
  max = [max, node.val].max
  min = [min, node.val].min
  [
    max_ancestor_diff_aux(node.left, max, min),
    max_ancestor_diff_aux(node.right, max, min)
  ].max
end

# def max_ancestor_diff(root)
#   left = root.left ? max_ancestor_diff(root.left) : 0
#   right = root.right ? max_ancestor_diff(root.right) : 0

#   p ['min', root.val, '-', min_child(root), (root.val - min_child(root)).abs]
#   p ['max', root.val, '-', max_child(root), (root.val - max_child(root)).abs]
#   r = [
#     (root.val - min_child(root)).abs,
#     (root.val - max_child(root)).abs
#   ].max
#   p ['r', r]
#   puts '------'
#   r
# end

# ===============================================
#       8
#     /    \
#   3       10
#  / \         \
# 1   6         14
#    / \      /
#   4   7   13
root = TreeNode.new(8)
root.left = TreeNode.new(3)
root.right = TreeNode.new(10)
root.left.left = TreeNode.new(1)
root.left.right = TreeNode.new(6)
root.left.right.left = TreeNode.new(4)
root.left.right.right = TreeNode.new(7)
root.right.right = TreeNode.new(14)
root.right.right.left = TreeNode.new(13)
# # root = [8,3,10,1,6,nil,14,nil,nil,4,7,13]
# # ===============================================

# # p min_child(root)
# # p max_child(root)
# p max_ancestor_diff(root)

# ===============================================
#        1
#          \
#           2
#             \
#              0
#             /
#            3
root = TreeNode.new(1)
root.right = TreeNode.new(2)
root.right.right = TreeNode.new(0)
root.right.right.left = TreeNode.new(3)
# ===============================================
# p min_child(root.right.right)
# p max_child(root.right.right)
p max_ancestor_diff(root)
