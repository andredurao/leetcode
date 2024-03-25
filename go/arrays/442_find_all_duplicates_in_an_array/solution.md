# Intuition
I need to store the visited numbers somehow in a hash, array or other. Hash is the best choice for lookup in this example because it can find the element in O(1). But storing `n` elements in a hash sounded like different than the problem challenge says *`"You must write an algorithm that runs in O(n) time and uses only constant extra space"`* constant extra space sounded for me like they were expecting me to solve it with space complexity `O(1)`. But I could not find a solution for that, I did with `O(n)`


# Approach
1. Start an empty hash with int keys, and since there could be only 1 or 2 elements of each, I'm not interested in the count, so their values could be empty structs for now
2. Start an empty array of ints to hold the solution
3. Iterate nums and check is present in the hash, if so concatenate it in the array, otherwise add it's value into the hash
4. Return the array(2)

# Complexity
- Time complexity: $$O(n)$$

- Space complexity: $$O(n)$$

# Code
```
func findDuplicates(nums []int) []int {
  res := []int{}
  numsMap := map[int]struct{}{}

  for _, num := range nums {
    _, found := numsMap[num]
    if found {
      res = append(res, num)
    } else {
      numsMap[num] = struct{}{}
    }
  }

  return res
}
```
