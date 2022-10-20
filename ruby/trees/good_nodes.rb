require './util.rb'
require './tree.rb'

def good_nodes(root)
  path_builder(root)
end

def path_builder(root, path=[])
  return 0 unless root
  current_path = path.dup << root.val
  l = path_builder(root.left, current_path)
  r = path_builder(root.right, current_path)
  current_val = current_path.max > root.val ? 0 : 1
  current_val + l + r
end

tree = [3,1,4,3,nil,1,5]
root = Util.build(tree)
p good_nodes(root)
