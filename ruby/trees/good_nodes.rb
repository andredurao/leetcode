require './util.rb'
require './tree.rb'

def good_nodes(root)
  path_builder(root)
end

def path_builder(root, max=nil)
  return 0 unless root
  max = max ? [max, root.val].max : root.val
  l = path_builder(root.left, max)
  r = path_builder(root.right, max)
  current_val = max > root.val ? 0 : 1
  current_val + l + r
end

tree = [3,1,4,3,nil,1,5]
root = Util.build(tree)
p good_nodes(root)
