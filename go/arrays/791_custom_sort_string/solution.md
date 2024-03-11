# Intuition
My first thought is that I needed to get the frequency of chars and rebuild the string using the custom order

# Approach
1. Get the char frequencies in a hash<char> -> int
2. Iterate the custom order and concatenate (n times) the char (ch) on a string buffer
2.1. Remove the ch in frequencies hashmap
3. Get remaining chars in the hashmap and sort them (remaining order)
4. Iterate the remaining order and concatenate (n times) the char (ch) on a string buffer
5. Return the string in buffer

# Complexity
- Time complexity: $$O(n * log(n))$$

- Space complexity: $$O(n)$$

# Code
```
func customSortString(order string, s string) string {
  var buffer bytes.Buffer
  charMap := map[rune]int{}

  // split chars and their frequencies
  for _, char := range s {
    _, found := charMap[char]
    if !found {
      charMap[char] = 0
    }
    charMap[char]++
  }

  // rebuild string in custom order
  for _, orderChar := range order {
    qty, found := charMap[orderChar]
    if found {
      for i:=0;i<qty;i++{
        buffer.WriteRune(orderChar)
      }
      delete(charMap, orderChar)
    }
  }

  // adds remaining chars
  remainingOrder := []int{}
  for key := range charMap {
    remainingOrder = append(remainingOrder, int(key))
  }
  sort.IntSlice.Sort(remainingOrder)

  for _, orderChar := range remainingOrder {
    qty, _ := charMap[rune(orderChar)]
    for i:=0;i<qty;i++{
      buffer.WriteRune(rune(orderChar))
    }
  }

  return buffer.String()
}
```
