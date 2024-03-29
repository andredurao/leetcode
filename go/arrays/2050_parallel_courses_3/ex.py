from collections import defaultdict, deque
from typing import List

class Solution:
    def minimumTime(self, n: int, relations: List[List[int]], time: List[int]) -> int:
        graph = defaultdict(list)
        indegree = [0] * n

        for (x, y) in relations:
            graph[x - 1].append(y - 1)
            indegree[y - 1] += 1

        queue = deque()
        max_time = [0] * n
        for node in range(n):
            if indegree[node] == 0:
                queue.append(node)
                max_time[node] = time[node]

        while queue:
            node = queue.popleft()
            print(max_time)
            print(graph[node])
            for neighbor in graph[node]:
                print("max", max(max_time[neighbor], max_time[node] + time[neighbor]))
                max_time[neighbor] = max(max_time[neighbor], max_time[node] + time[neighbor])
                indegree[neighbor] -= 1
                if indegree[neighbor] == 0:
                    queue.append(neighbor)

        return max(max_time)



n = 5
relations = [[1,5],[2,5],[3,5],[3,4],[4,5]]
t = [1,2,3,4,5]
s = Solution()
result = s.minimumTime(n, relations, t)

print(result)
