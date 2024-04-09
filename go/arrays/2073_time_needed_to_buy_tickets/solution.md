# Intuition
I've tried to simulate the behaviour explained in the description, but then I've realized that removing elements from the array could take longer than having a condition to check when elements reach zero.

# Approach
1. Initialize a var `total`
2. Repeat the following in an infinite loop
3. Iterate the tickets array and check if the current element value is greater than zero, if so decrement the element and increment total
4. If the element in position k reached zero, return the total

# Complexity
- Time complexity: $$O(n^2)$$

- Space complexity: $$O(1)$$

# Code
```
func timeRequiredToBuy(tickets []int, k int) int {
    total := 0

    for {
        for i := range tickets{
            if tickets[i] > 0 {
                tickets[i]--
                total++
            }
            if i == k && tickets[i] == 0 {
                return total
            }
        }
    }
}
```
