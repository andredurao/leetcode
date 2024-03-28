# Intuition
I knew that I've had to save the frequencies somehow in a hashmap or array, but I was not sure about the subarrays. The good thing is that in this problem they were not interested in subsets of the existing subarray, but the largest subarray that matches the condition "good subarray".

# Approach
1. Initialize two cursors (`l` and `r`) to search for subarrays as zero, and also a `length` variable to store the max length found so far. Finally a hashmap `frequencyMap` to store the frequencies
2. Repeat the following until `r` reaches the size of `nums`:
2.1. let `num` = `nums[r]`
2.2. increment `frequency[num]`
2.2. if `frequency[num]` is greater than `k`, we need to get `l` closer to `r` while decrementing all frequencies found in the way, until `frequency[num]` is still greater than `k`. In other words, close the window (get it narrowed).
3. return the maximum of either `length` or `r - l`, because `length` is only updated when a value exceeds `k`, and with that we'll get the true value even if no value has exceeded `k`

# Complexity
- Time complexity: $$O(n)$$

- Space complexity: $$O(n)$$

# Code
```
func maxSubarrayLength(nums []int, k int) int {
  length := 0
  l := 0
  r := 0
  num := 0
  frequencyMap := map[int]int{}

  for r < len(nums) {
    num = nums[r]
    _, found := frequencyMap[num]
    if !found {
      frequencyMap[num] = 0
    }
    frequencyMap[num]++

    if frequencyMap[num] > k {
      length = max(length, r-l)
      // slide left until the current value has been found
      for frequencyMap[num] > k {
        frequencyMap[nums[l]]--
        l++
      }
    }
    r++
  }

  return max(length, r-l)
}
```
