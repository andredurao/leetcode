# Intuition
On this problem I knew that I needed to traverse list1 until points a and b and save them to cut off the nodes between those.

Is also needed to traverse the entire list2 because there's no reference for `tail`

# Approach
1. Save the head of list1 (`first`)
2. Traverse list 1 while incrementing a var (`i`) and save the references for nodes a (*one node before a*) and b (*one node after b*). Stop on node b.
3. Set the a.next to the head of list2
4. Traverse list2 until it's last node
5. Set the next node as b, the remaining of the list1 after b

# Complexity
- Time complexity: $$O(m+n)$$

- Space complexity: $$O(1)$$

# Code
```
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
  first := list1
  i := 0
  aNode := &ListNode{}
  bNode := &ListNode{}
  // [0 1 2 3 4 5 6 7] 3, 4
  for i <= b {
    if i == a - 1 {
      aNode = list1
    }
    list1 = list1.Next
    i++
  }
  bNode = list1

  aNode.Next = list2

  for aNode.Next != nil {
    aNode = aNode.Next
  }

  aNode.Next = bNode

  return first
}
```
