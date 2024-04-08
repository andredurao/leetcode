// https://leetcode.com/problems/number-of-students-unable-to-eat-lunch/description/

package main

import "fmt"

func main() {
	students := []int{1, 1, 0, 0}
	sandwiches := []int{0, 1, 0, 1}
	fmt.Println(countStudents(students, sandwiches))
}

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
