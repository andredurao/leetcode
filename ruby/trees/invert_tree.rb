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
# @return {TreeNode}
def invert_tree(root)
  if root
    root.left, root.right = root.right, root.left
    invert_tree(root.left) if root.left
    invert_tree(root.right) if root.right
  end
  root
end

def in_order_traversal(node, items=[])
  in_order_traversal(node.left, items) if node.left
  items << node.val
  in_order_traversal(node.right, items) if node.right
  items
end

# ============================
#       4
#   2       7
# 1   3   6   9
root = TreeNode.new(4)
root.left = TreeNode.new(2)
root.right = TreeNode.new(7)
root.left.left = TreeNode.new(1)
root.left.right = TreeNode.new(3)
root.right.left = TreeNode.new(6)
root.right.right = TreeNode.new(9)
# ============================

p in_order_traversal(root)

invert_tree(root)

p in_order_traversal(root)

