class Node
  attr_accessor :left, :right, :val

  def initialize(val)
    @val = val
  end
end

def inorder_traversal(root)
  return nil unless root
  traversal_helper(root)
end

def traversal_helper(node, items=[])
  return [] unless node
  traversal_helper(node.left, items)
  items << node.value
  traversal_helper(node.right, items)
  items
end

# ============================
#       4
#   2       6
# 1   3   5   7
root = Node.new(4)
root.left = Node.new(2)
root.right = Node.new(6)
root.left.left = Node.new(1)
root.left.right = Node.new(3)
root.right.left = Node.new(5)
root.right.right = Node.new(7)
# ============================

p 'inorder_traversal'
p inorder_traversal(Node.new(nil))
# p inorder_traversal(nil)
