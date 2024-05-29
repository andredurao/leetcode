# Intuition
Initially, I've tried to simply convert to an int value only to realize later that the size constraint was from 1 to 500 bits ;(

# Approach
1. Convert the string to a array of byte size values (`val`)
2. Initialize a counter variable
3. Until the length of `val` is different than `'1'`
3.1. let `end` be the last char of `val`
3.2. if `end` equals to `0` shift val right 1 time and increment the counter
3.3. otherwise lookback from `end` to 0 until a `0` value is found while incrementing the counter, assign that position to a var `j`
3.3.1. Increment counter
3.3.2. if `j` equals to 0, increment counter and set `val` to `'1'`
3.3.3. otherwise set `val[j]` to `1` and cut the right side of val after j

# Complexity
- Time complexity: $$O(m*n)$$

- Space complexity: $$O(n)$$

# Code
```
func numSteps(s string) int {
	val := []byte(s)
	steps := 0

	for len(val) > 0 {
		// fmt.Println("S", string(val), len(val), steps)
		if len(val) == 1 && val[0] == '1' {
			break
		}
		end := len(val) - 1

		if val[end] == '0' {
			val = val[:end] // SHR
			steps++
		} else {
			j := end
			for j > 0 && val[j] == '1' {
				steps++
				j--
			}
			steps++
			if j == 0 {
				steps++
				val = []byte("1")
			} else {
				val[j] = '1'
				val = val[0 : j+1]
			}
		}
		// fmt.Println("E", string(val), len(val), steps, "\n")
	}

	return steps
}

```
