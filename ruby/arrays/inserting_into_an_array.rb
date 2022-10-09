# @param {Integer[]} nums1
# @param {Integer} m
# @param {Integer[]} nums2
# @param {Integer} n
# @return {Void} Do not return anything, modify nums1 in-place instead.
# Complexity analysis: Time O(n+m) Space O(n+m)
def merge(nums1, m, nums2, n)
  result = []
  m_index = n_index = 0
  while m_index < m || n_index < n
    # conditions:
    # drain OR smaller than the other queue
    if m_index == m || (n_index < n && n > 0 && nums2[n_index] <= nums1[m_index])
      result.push nums2[n_index]
      n_index += 1
    else n_index == n || (m_index < m && m > 0 && nums1[m_index] <=  nums2[n_index])
      result.push nums1[m_index]
      m_index += 1
    end
  end

  m_index = -1
  nums1.map! do |val|
    m_index += 1
    result[m_index]
  end
end

# Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
# Output: [1,2,2,3,5,6]
nums1 = [1,2,3,0,0,0]
m = 3
nums2 = [2,5,6]
n = 3

merge(nums1, m, nums2, n)
p nums1
