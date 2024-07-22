//

package main

import (
	"fmt"
	"sort"
)

func main() {
	// names := []string{"Mary", "John", "Emma"}
	// heights := []int{180, 165, 170}

	names := []string{"Alice", "Bob", "Bob"}
	heights := []int{155, 185, 150}

	fmt.Println(sortPeople(names, heights))
}

type Person struct {
	Name   string
	Height int
}

func sortPeople(names []string, heights []int) []string {
	// zip people in a single struct
	people := make([]Person, len(names))
	for i := range people {
		people[i].Name = names[i]
		people[i].Height = heights[i]
	}

	sort.Slice(people, func(a, b int) bool { return people[b].Height < people[a].Height })

	for i := range people {
		names[i] = people[i].Name
	}
	return names
}
