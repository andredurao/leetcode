# Intuition
I've read the description, but this part I've got confused with this part here: *`"You must implement an algorithm that runs in O(n) time and uses O(1) auxiliary space."`*, because I thought that it was required to solve it in O(1) space, but that was `auxiliary space`. After realizing that I could still use `O(n)` space, I've realized that I could use a hash to mark each positive integer found (O(n)) and the max positive value (aux O(1)) and later iterate from 1 to the max positive value, looking up in the hash for missing values.

# Approach
1. Initialize a hash table using int keys and maxValue as zero
2. Iterate the nums array and if a num is positive do the following
2.1. mark that number as visited
2.2. check if that number is greater than maxValue, if so update maxValue
3. iterate from 1 (minimum positive number) to maxValue and if a value is not present in the hash table return it, otherwise return `maxValue + 1`

# Complexity
- Time complexity: $$O(n)$$

- Space complexity: $$O(n)$$

# Code
```

func firstMissingPositive(nums []int) int {
  positivesMap := map[int]struct{}{}
  maxNum := 0

  for _, num := range nums {
    if num > 0 {
      maxNum = max(maxNum, num)
      positivesMap[num] = struct{}{}
    }
  }


  for i := 1; i < maxNum ; i++ {
    _, found := positivesMap[i]
    if !found {
      return i
    }
  }

  return maxNum+1
}
```
