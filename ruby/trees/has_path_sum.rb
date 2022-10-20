require './util.rb'
require './tree.rb'

# @param {TreeNode} root
# @param {Integer} target_sum
# @return {Boolean}
def has_path_sum(root, target_sum)
  return false unless root
  target_sum -= root.val
  return true if target_sum == 0 && !root.left && !root.right
  has_path_sum(root.left, target_sum) || has_path_sum(root.right, target_sum)
end


tree = [1,2,3]
target_sum = 1
root = Util.build(tree)
p has_path_sum(root, target_sum)

tree = [1,2,3]
target_sum = 3
root = Util.build(tree)
p has_path_sum(root, target_sum)

tree = [5,4,8,11,nil,13,4,7,2,nil,nil,nil,1]
target_sum = 22
root = Util.build(tree)
p has_path_sum(root, target_sum)
