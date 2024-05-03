# https://leetcode.com/problems/compare-version-numbers/

# @param {String} version1
# @param {String} version2
# @return {Integer}
def compare_version(version1, version2)
  v1_revisions = version1.split('.').map(&:to_i)
  v2_revisions = version2.split('.').map(&:to_i)

  max_revisions_size = [v1_revisions.size, v2_revisions.size].max

  (0..max_revisions_size).each do |i|
    v1 = v1_revisions[i] || 0
    v2 = v2_revisions[i] || 0
    return 1 if v1 > v2
    return -1 if v2 > v1
  end

  return 0
end

v1 = "1.01"
v2 = "1.001"
p compare_version(version1, version2)
