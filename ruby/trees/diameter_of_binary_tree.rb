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
def diameter_of_binary_tree(root)
  leaves = build_leaves(root)
  paths = leaves.reduce({}) do |hash, leaf|
    hash[leaf] = build_path(root, leaf)
    hash
  end

  if leaves.size == 1
    return paths[leaves.first].size - 1
  else
    # bruteforce lookup for all pairs of leaves if there's more than one leaf
    max = 0
    leaves.permutation(2).each do |permutation|
      path_a = paths[permutation[0]]
      path_b = paths[permutation[1]]

      node_in_common = (path_a | path_b).last
      path_a.index

      # this path goes all the way from a to b
      max_path = [
        (path_a - path_b | path_b - path_a).size,
        path_a.size - 1,
        path_b.size - 1
      ].max
      max = max_path if max_path > max
    end
    max
  end
end

# @return [Array[TreeNode]] the collection of leaves in this tree
def build_leaves(root, result=[])
  # p [root, result]
  return [root] if !root.left && !root.right
  result += build_leaves(root.left, result) if root.left
  result += build_leaves(root.right, result) if root.right
  result
end

def build_path(root, node, path=[])
  return unless root
  path = path.dup
  # p [root.val, node.val, path.map(&:val)]
  path << root
  return path if root.val == node.val

  build_path(root.left, node, path) || build_path(root.right, node, path)
end


# ===============================================
#       1
#     /    \
#   2       3
#  / \
# 4   5
# root = TreeNode.new(1)
# root.left = TreeNode.new(2)
# root.right = TreeNode.new(3)
# root.left.left = TreeNode.new(4)
# root.left.right = TreeNode.new(5)
# p diameter_of_binary_tree(root)

# ===============================================
#       1
#     /
#   2
# root = TreeNode.new(1)
# root.left = TreeNode.new(2)
# p diameter_of_binary_tree(root)

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
