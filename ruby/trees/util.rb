def insert(node, val)
  dir = val > node.val ? :right : :left
  if node.send(dir)
    insert(node.send(dir), val)
  else
    node.send("#{dir}=", TreeNode.new(val))
  end
end

def populate(array)
  root = TreeNode.new(array.shift)
  array.each{|val| insert(node, val) if val}
  root
end

def in_order_traversal(node, items=[])
  in_order_traversal(node.left, items) if node.left
  items << node.val
  in_order_traversal(node.right, items) if node.right
  items
end

def min_child(node)
  min_left  = node.left  ? min_child(node.left)  : node.val
  min_right = node.right ? min_child(node.right) : node.val
  [min_left, min_right].min
end

def max_child(node)
  max_left  = node.left  ? max_child(node.left)  : node.val
  max_right = node.right ? max_child(node.right) : node.val
  [max_left, max_right].max
end
