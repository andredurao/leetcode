class Node
  attr_accessor :left, :right, :value

  def initialize(value)
    @value = value
  end
end

def pre_traversal(node, items=[])
  items << node.value
  pre_traversal(node.left, items) if node.left
  pre_traversal(node.right, items) if node.right
  return items
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

p 'pre_traversal'
p pre_traversal(root)
