# https://leetcode.com/problems/find-the-town-judge/

# @param {Integer} n
# @param {Integer[][]} trust
# @return {Integer}
def find_judge(n, trust)
  trusted_map = {}
  truster_map = {}
  trust.each do |truster, trusted|
    trusted_map[trusted] ||= []
    trusted_map[trusted] << truster

    truster_map[truster] ||= []
    truster_map[truster] << trusted
  end

  judge = -1

  trusted_map.each do |trusted, people|
    judge = trusted if people.size == n - 1 && !truster_map[trusted]
  end

  judge
end


n = 2
trust = [[1,2]]
p find_judge(n, trust)


n = 3
trust = [[1,3],[2,3]]
p find_judge(n, trust)


n = 3
trust = [[1,3],[2,3],[3,1]]
p find_judge(n, trust)


n = 1
trust = []
p find_judge(n, trust)
