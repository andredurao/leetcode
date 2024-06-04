# Intuition
The first time I read the description I thought that this problem was related to find the longest substring that was also a palindrome, but in reality they were expecting the longest string that can be formed with the characters included in the string that was provided.

# Approach
1. Create a hash map to count the frequencies of chars contained in the input
2. Set total to length of the input
3. A palindrome could have only one substring of odd elements, the ones contained in the middle of the string. To ensure that the generated string would return the right length, it was needed the count of odd elements and reduce it until we have at maximum only one odd substring

# Complexity
- Time complexity: $$O(n)$$

- Space complexity: $$O(n)$$

# Code
```
func longestPalindrome(s string) int {
    f := map[rune]int{}

    for _, ch := range s {
        if _, found := f[ch] ; !found {
            f[ch] = 0
        }
        f[ch]++
    }

    total := len(s)

    odds := 0
    for _, count := range f {
        if count % 2 == 1 {
            odds++
        }
    }

    for odds > 1 {
        total--
        odds--
    }

    return total
}
```
