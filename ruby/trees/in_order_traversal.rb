class Node
  attr_accessor :left, :right, :value

  def initialize(value)
    @value = value
  end
end

def in_order_traversal(node, items=[])
  in_order_traversal(node.left, items) if node.left
  items << node.value
  in_order_traversal(node.right, items) if node.right
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

p 'in_order_traversal'
p in_order_traversal(root)
