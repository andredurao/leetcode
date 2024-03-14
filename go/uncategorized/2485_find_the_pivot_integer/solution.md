# Intuition
My first thought was to brute force the search and that would take $$O(n)$$, but then I changed my idea to lookup for values using binary search.

|     n     |  1 |  2 |  3 |  4 |  5 |  **6** | 7  |  8 |
|:---------:|:--:|:--:|:--:|:--:|:--:|:--:|----|:--:|
|  left sum |  1 |  3 |  6 | 10 | 15 | **21** | 28 | 36 |
| right sum | 36 | 35 | 33 | 30 | 26 | **21** | 15 |  8 |

* Σ = $$(n * (n-1)) / 2 | n = 8$$
* Left sum = $$(n * (n-1)) / 2$$
* Right sum = $$Σ - (LeftSum (n-2)) / 2$$

# Approach
1. Initialize sigma as the total sum of 1..n, l as 1 and r as 1
2. Set mid as the average of L and R, and use a binary search to lookup for a mid point on which left sum and right sum are equal

# Complexity
- Time complexity: $$O(n *log(n))$$

- Space complexity: $$O(1)$$

# Code
```
func pivotInteger(n int) int {
  sigma := (n*(n+1))/2
  l := 1
  r := n

  for l <= r {
    mid := (l+r) / 2
    lsum := (mid*(mid+1))/2
    rsum := sigma - ((mid-1)*mid)/2
    if lsum == rsum {
      return mid
    } else if lsum < rsum {
      l = mid + 1
    } else {
      r = mid - 1
    }
  }

  return -1
}
```
