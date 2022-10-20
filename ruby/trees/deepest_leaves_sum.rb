require './util.rb'
require './tree.rb'

# @param {TreeNode} root
# @return {Integer}
def deepest_leaves_sum(root)
  map = {}
  max_height = group_vals_by_height(root, map)
  map[max_height]
  # A BFS performs better than recursion in this case (mc kinley)
end

def group_vals_by_height(root, map, height=1)
  return 0 unless root
  map[height] = map[height] ? (map[height] + root.val) : root.val
  l = group_vals_by_height(root.left, map, height+1)
  r = group_vals_by_height(root.right, map, height+1)
  [l+1, r+1].max
end

tree = [1,2,3,4,5,nil,6,7,nil,nil,nil,nil,8]
root = Util.build(tree)
p deepest_leaves_sum(root)
