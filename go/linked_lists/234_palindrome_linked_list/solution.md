# Intuition
My first thought is that I need to save all values in a list and check if that list is a palindrome

# Approach
1. Iterate the linked list while appending each value in a array
2. Iterate from the first position (i) to the middle of the list and check if each value is the same in `values[i]` and it's opposite end position `values[len(values)-1-i]`; Otherwise return false
3. If the end was reached, then the linked list is a palindrome

# Complexity
- Time complexity: $$O(n)$$

- Space complexity: $$O(n)$$


# Code
```
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
  values := []int{}

  for head != nil {
    values = append(values, head.Val)
    head = head.Next
  }

  for i := 0 ; i < len(values) / 2 ; i++ {
    if values[i] != values[len(values)-i-1] {
      return false
    }
  }

  return true
}
```
