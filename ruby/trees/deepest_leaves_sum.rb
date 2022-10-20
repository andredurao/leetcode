require './util.rb'
require './tree.rb'

# @param {TreeNode} root
# @return {Integer}
def deepest_leaves_sum(root)
  heights_map = group_vals_by_height(root)
  max = heights_map.keys.max
  heights_map[max].reduce(:+)
end

def group_vals_by_height(root, map={}, height=0)
  return unless root
  map[height] ||= []
  map[height] << root.val
  group_vals_by_height(root.left, map, height+1)
  group_vals_by_height(root.right, map, height+1)
  map
end

tree = [1,2,3,4,5,nil,6,7,nil,nil,nil,nil,8]
root = Util.build(tree)
p deepest_leaves_sum(root)
