# Intuition
I knew that I this could be done in many different ways, but since we are interested in the last word. I thought it could be done starting from the end of the string and increment a value until the value of the cursor is different than space, but example 2 had a string that ended with spaces, to satisfy that condition I've moved the cursor to the last char that was not a space before counting

# Approach
1. Initialize a cursor `i` and a total `res`, both with `0`
2. Move the cursor, from right to left, until it stops at the first non-space character
3. Move the cursor, from right to left, while incrementing `res` until it reaches zero. If the character under `cursor` is a space, return `res`
4. Return `res`

# Complexity
- Time complexity: $$O(n)$$

- Space complexity: $$O(1)$$

# Code
```
func lengthOfLastWord(s string) int {
    res := 0
    i := len(s) - 1

    for s[i] == ' ' {
        i--
    }

    for ; i >= 0 ; i-- {
        if s[i] == ' ' {
            return res
        }
        res++
    }
    return res
}
```
