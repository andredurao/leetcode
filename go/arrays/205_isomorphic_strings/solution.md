# Intuition
For this problem I've knew that I've had to store the mapped values and checked for the first time they show up, and later check if the values were being used correctly

# Approach
Initially I thought that I had to map and test the values from `s` to `t` but later one test case failed because it was mapping a different value to that was already mapped on `t`. With that I've used two maps, one from `s` to `t` and one from `t` to `s`

1. Initialize two byte arrays to store the maps from s->t and t->s
2. Use a cursor `i` and iterate from 0 to length of `s` (s and t are the same size)
3. Check if `s[i] != t[i]`, if there is a char mapped on `s[i]` otherwise store both references `s[i] -> t[i]` vice versa, unless that there is a `t[i]` already mapped to a different value

# Complexity
- Time complexity: $$O(n)$$
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity: $$O(1)$$


# Code
```
func isIsomorphic(s string, t string) bool {
    sCharmap := make([]rune, 256)
    tCharmap := make([]rune, 256)

    for i := range s {
        if sCharmap[rune(s[i])] == 0 && tCharmap[rune(t[i])] == 0 {
            sCharmap[rune(s[i])] = rune(t[i])
            tCharmap[rune(t[i])] = rune(s[i])
        } else {
            if sCharmap[rune(s[i])] != rune(t[i]) {
                return false
            }
        }
    }

    return true
}
```
