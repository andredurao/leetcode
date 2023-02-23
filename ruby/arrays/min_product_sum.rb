# https://leetcode.com/problems/minimize-product-sum-of-two-arrays/submissions/903562941/
def ranks(array)
  array_with_index = array.dup.each.with_index.to_a
  result = []
  array_with_index.sort.each_with_index{|val, i| result[val[1]] = i} # rank from their indexess
  result
end

# @param {Integer[]} nums1
# @param {Integer[]} nums2
# @return {Integer}
def min_product_sum(nums1, nums2)
  size = nums1.size
  rankings = ranks(nums2)
  grouped_by_reverse_nums2 = []

  nums1.sort!
  nums2.each_with_index do |num2, index|
    ranking = rankings[index]
    # p [num2, ranking, nums1[size - 1 - ranking]]
    grouped_by_reverse_nums2 << nums1[size - 1 - ranking]
  end

  total = 0
  grouped_by_reverse_nums2.zip(nums2) do |n1, n2|
    total += n1 * n2
  end
  total
end

nums1 = [5,3,4,2]
nums2 = [4,2,2,5]
p min_product_sum(nums1, nums2)

nums1 = [5,5,5,5,5]
nums2 = [5,5,5,5,5]
p min_product_sum(nums1, nums2)
