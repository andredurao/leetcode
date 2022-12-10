# https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/
require './util.rb'
require './tree.rb'

# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val = 0, left = nil, right = nil)
#         @val = val
#         @left = left
#         @right = right
#     end
# end
# @param {TreeNode} root
# @return {Integer}
def max_product(root)
  list = []
  total_sum = sum(root, list)

  result = 0
  list.each do |sum|
    temp_sum = total_sum - sum
    product = sum * temp_sum
    result = [result, product].max
  end
  result % 1000000007
end

def sum(root, list)
  return 0 unless root
  left_sum = sum(root.left, list)
  right_sum = sum(root.right, list)
  subtree_sum = left_sum + right_sum + root.val
  list << subtree_sum
  subtree_sum
end

tree = [1,2,3,4,5,6]
root = Util.build(tree)
list = []
p max_product(root)
