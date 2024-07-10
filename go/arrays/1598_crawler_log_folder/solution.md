# Intuition
I've used a simple to check if the dir name (items in the array) are different than "." or "..", if so I've assumed that was a directory name

# Approach
1. Initialize an integer variable `res`
2. Iterate through the logs array
2.1. If the current item refers to the parent dir ("../") decrease `res`, only if `res` is larger than zero
2.2. Otherwise increase res if the current item is different than `"./"`


# Complexity
- Time complexity: $$O(n)$$

- Space complexity: $$O(1)$$

# Code
```
func minOperations(logs []string) int {
	res := 0
	for _, dir := range logs {
		if dir == "../"{
			if res > 0 {
				res--
			}
		} else if dir != "./" {
			res++
		}
	}

	return res
}
```
