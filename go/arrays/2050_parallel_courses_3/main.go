// https://leetcode.com/problems/parallel-courses-iii/

package main

import (
	"fmt"
)

func main() {
	// n := 3
	// relations := [][]int{{1, 3}, {2, 3}}
	// time := []int{3, 2, 5}

	n := 5
	relations := [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}}
	time := []int{1, 2, 3, 4, 5}
	result := minimumTime(n, relations, time)
	fmt.Println(result)
}

type Course struct {
	Prerequisites []int
	Time          int
	Completed     bool
}

type MinCoursesResponses struct {
	Indexes  []int
	MinTime  int
	MinIndex int
}

func minimumTime(n int, relations [][]int, time []int) int {
	// build courses array
	courses := make([]Course, n)

	for i := 0; i < n; i++ {
		courses[i].Time = time[i]
		courses[i].Completed = false
		courses[i].Prerequisites = make([]int, 0)
	}

	for _, relation := range relations {
		index := relation[1] - 1
		courses[index].Prerequisites = append(courses[index].Prerequisites, relation[0]-1)
	}

	fmt.Printf("%v\n", courses)

	timeSpent := 0

	for {
		minCourses := coursesThatCanBeStarted(courses)
		fmt.Printf("%v\n", minCourses)

		if len(minCourses.Indexes) == 0 {
			return timeSpent
		}
		timeSpent = timeSpent + minCourses.MinTime

		for _, i := range minCourses.Indexes {
			courses[i].Time = courses[i].Time - minCourses.MinTime
			if courses[i].Time <= 0 {
				courses[i].Completed = true
			}
		}

		fmt.Printf("%v\n", courses)
	}
}

func coursesThatCanBeStarted(courses []Course) (result MinCoursesResponses) {
	for index, course := range courses {
		if !course.Completed {
			completedPrerequisites := 0
			for _, prerequisite := range course.Prerequisites {
				if courses[prerequisite].Completed {
					completedPrerequisites++
				}
			}
			if completedPrerequisites == len(course.Prerequisites) {
				if len(result.Indexes) == 0 {
					result.MinTime = course.Time
					result.MinIndex = index
				}
				result.Indexes = append(result.Indexes, index)

				if course.Time < result.MinTime {
					result.MinTime = course.Time
					result.MinIndex = index
				}
			}
		}
	}
	return
}
