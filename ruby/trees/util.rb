class Util
  def self.build(input)
    return if input == []
    nodes = input.map do |val|
      TreeNode.new(val) if val
    end
    repo = nodes.dup
    root = repo.shift

    nodes.each do |node|
      next unless node
      node.left = repo.shift
      node.right = repo.shift
    end
    root
  end

  def self.insert(node, val)
    dir = val > node.val ? :right : :left
    if node.send(dir)
      insert(node.send(dir), val)
    else
      node.send("#{dir}=", TreeNode.new(val))
    end
  end

  def self.populate(array)
    root = TreeNode.new(array.shift)
    array.each{|val| insert(node, val) if val}
    root
  end

  def self.in_order_traversal(node, items=[])
    in_order_traversal(node.left, items) if node.left
    items << node.val
    in_order_traversal(node.right, items) if node.right
    items
  end

  def self.min_child(node)
    min_left  = node.left  ? min_child(node.left)  : node.val
    min_right = node.right ? min_child(node.right) : node.val
    [min_left, min_right].min
  end

  def self.max_child(node)
    max_left  = node.left  ? max_child(node.left)  : node.val
    max_right = node.right ? max_child(node.right) : node.val
    [max_left, max_right].max
  end

  def self.max(node)
    return -Float::INFINITY unless node
    l = max(node.left)
    r = max(node.right)
    [l, r, node.val].max
  end
end
