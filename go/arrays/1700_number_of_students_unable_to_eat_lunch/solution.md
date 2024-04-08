# Intuition
From the problem description was clear that it was necessary to pop the head element from student's list, check if it suits the head of sandwiches list and they don't match, append them back to the end of their list. What was not clear is how to break this loop, or how to check if all students are different than the sandwich head element, if I choose to check entire list of students on each iteration that would increase the cost of the solution by 1 order. So I chosen to increment a counter on each time there was no match and reset it during matches, with that when the counter reaches the size of the list of students I knew that it was over

# Approach
1. Initialize a counter to count how many times a student went back to the end of the line consecutively `backToQueueEnd`
2. Repeat until the counter is smaller that the size of students and there are still students in the list:
2.1. Reset the counter If the head elements of both lists are equal and remove the head elements of both lists, otherwise increment the counter, remove the head element of students and append it to the end of the list

# Complexity
- Time complexity: $$O(n^2)$$


- Space complexity: $$O(1)$$

# Code
```
func countStudents(students []int, sandwiches []int) int {
    backToQueueEnd := 0
    for backToQueueEnd < len(students) && len(students) > 0 {
        // fmt.Printf("stu %v sand %v backCount %d\n", students, sandwiches, backToQueueEnd)
        if students[0] == sandwiches[0] {
            // fmt.Println("consume")
            backToQueueEnd = 0
            students = students[1:]
            sandwiches = sandwiches[1:]
        } else {
            // fmt.Println("back")
            backToQueueEnd++
            students = append(students[1:], students[0])
        }
    }
    return len(students)
}
```
