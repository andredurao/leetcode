# Intuition
Initially I've thought that it was necessary to use a sliding window, but I wrote a not so clever and simple solution

# Approach
1. Early return false in case $$arr$$ has less then 3 items
2. Iterate from 0 to length of $$arr - 3$$ and check if the 3 numbers on the left (starting from i) are odds, if so increment count, otherwise save the index on a var ($$lastEven$$). At the end return true if count is equal to 3, otherwise repeat it with $$i = lastEven+1$$

# Complexity
- Time complexity: $$O(n)$$

- Space complexity: $$O(1)$$

# Code
```
func threeConsecutiveOdds(arr []int) bool {
    if len(arr) < 3 {
        return false
    }
    count, i, lastEven := 0, 0, -1
    for i <= len(arr) - 3 {
        count = 0
        if arr[i] % 2 == 1 {
            count++
        } else {
            lastEven = i
        }
        if arr[i+1] % 2 == 1 {
            count++
        } else {
            lastEven = i+1
        }
        if arr[i+2] % 2 == 1 {
            count++
        } else {
            lastEven = i+2
        }

        if count == 3 {
            return true
        }
        i = lastEven+1
    }
    return false
}
```
