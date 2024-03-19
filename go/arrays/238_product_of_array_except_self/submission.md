# Intuition
I knew that I could store a product total in a variable and fill up the expected array by iterating the original `nums` and dividing total by `num[n]` on each position to give the results, but I forgot to double check for zero values and let a division by zero to take place on example 2. To solve that I use two other counters:
1. Total product without zero: to use in case `nums[n]` equals to zero
2. Zero count: to check if there's more than one zero in `nums` and act accordingly on each division

# Approach
1. Start the counters and product variables
1.1. totalProduct = 1
1.2. totalProductWithoutZero = 1
1.3. zeroCount = 0
2. Iterate through nums and to the following on each step:
2.1. Multiply the total product
2.2. Increase zero count, on each zero found
2.3. Multiply the zero count
3. If there's more than one zero, the total product would always be zero
4. Iterate again through nums and now fill the result array dividing the total T by `nums[n]`, given that T can be one of the following:
4.1. totalProductWithoutZero when `nums[n]` = 0
4.2. totalProduct / num when `nums[n]` ≠ 0

# Complexity
- Time complexity: $$O(n)$$

- Space complexity: $$O(n)$$

# Code
```
func productExceptSelf(nums []int) []int {
  totalProduct := 1
  totalProductWithoutZero := 1
  zeroCount := 0
  for _, num := range nums {
    totalProduct *= num
    if num == 0 {
      zeroCount++
    } else {
      totalProductWithoutZero *= num
    }
  }
  if zeroCount > 1 {
    totalProductWithoutZero = 0
  }
  res := make([]int, len(nums))
  for i, num := range nums {
    if num == 0 {
      res[i] = totalProductWithoutZero
    } else {
      res[i] = totalProduct / num
    }
  }
  return res
}
```
