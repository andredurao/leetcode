# Intuition

Having two cursors for L and R checking for repeated prefixes and suffixes

# Approach

1. Start L with 0 and R with the last position of the string
2. While chars on L and R positions are the same do the following
2.1. Save the char on the left as `ch`
2.2. Move L to the right until `s[l]` = ch and L < R
2.3. Move R to the left until `s[r]` = ch and R > L
2.4. Cut off the prefix and suffix before L and after R

# Complexity
- Time complexity: $$O(n)$$

- Space complexity: $$O(1)$$

# Code
```
func minimumLength(s string) int {
	l := 0
	r := len(s) - 1
	for len(s) > 0 && s[l] == s[r] && l < r {
		ch := s[l]
		for l < (r-1) && ch == s[l+1] {
			l++
		}
		for r > (l+1) && ch == s[r-1] {
			r--
		}
		// cut off prefix and suffix

		s = s[(l + 1):r]
		l = 0
		r = len(s) - 1
	}

	return len(s)
}

```
