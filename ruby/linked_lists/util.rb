class Util
  def self.deserialize(list)
    node = ListNode.new(list.shift)
    _next = nil
    list.each do |val|
      _next = node unless _next
      _next.next = ListNode.new(val)
      _next = _next.next
    end

    node
  end

  def self.serialize(node)
    result = []
    while node
      p node
      result << node.val
      node = node.next
    end
    result
  end
end
