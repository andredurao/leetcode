# Intuition

Somehow I have to save a hash containing the unique values and their quantities

# Approach

1. Save a hash containing the unique values and their frequencies
1.1. Get the max frequency in that hash (maxF)
2. Find the sum of frequencies (F) in the hash on which f = maxF

# Complexity
- Time complexity: $$O(n)$$

- Space complexity: $$O(n)$$

# Code
```
func maxFrequencyElements(nums []int) int {
	frequencyMap := map[int]int{}
	maxFreq := 0

	for _, num := range nums {
		_, found := frequencyMap[num]
		if !found {
			frequencyMap[num] = 0
		}
		frequencyMap[num]++
		maxFreq = int(max(float64(maxFreq), float64(frequencyMap[num])))
	}

	res := 0
	for _, freq := range frequencyMap {
		if freq == maxFreq {
			res += freq
		}
	}

	return res
}
```
