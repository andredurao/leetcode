package main

import (
	"fmt"
)

func main() {
	edges := [][]int{{0, 2, 5}, {0, 1, 2}, {1, 2, 1}, {3, 0, 3}}
	g := Constructor(4, edges)
	fmt.Println(g.ShortestPath(3, 2)) //6
	fmt.Println(g.ShortestPath(0, 3)) //-1
	g.AddEdge([]int{1, 3, 4})
	fmt.Println(g.ShortestPath(0, 3)) //6
}

type Graph struct {
	Edges      [][]int
	Vertices   map[int]bool
	EdgesMap   map[string]int
	Neighbours map[int]map[int]struct{}
}

func Constructor(n int, edges [][]int) Graph {
	edgesMap := make(map[string]int, 0)
	neighboursMap := make(map[int]map[int]struct{}, 0)
	vertices := make(map[int]bool, 0)
	graph := &Graph{edges, vertices, edgesMap, neighboursMap}
	for _, edge := range edges {
		graph.AddEdge(edge)
	}
	return *graph
}

func (this *Graph) AddEdge(edge []int) {
	this.Edges = append(this.Edges, edge)

	this.EdgesMap[fmt.Sprintf("%d-%d", edge[0], edge[1])] = edge[2]
	this.Vertices[edge[0]] = true
	this.Vertices[edge[1]] = true

	_, found := this.Neighbours[edge[0]]
	if !found {
		this.Neighbours[edge[0]] = make(map[int]struct{}, 0)
	}
	this.Neighbours[edge[0]][edge[1]] = struct{}{}
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	unvisited := make(map[int]bool, 0)
	distances := make(map[int]int, 0)
	for v := range this.Vertices {
		unvisited[v] = true
		distances[v] = int(1e7)
	}
	distances[node1] = 0
	current := node1

	for len(unvisited) > 0 && current != node2 {
		delete(unvisited, current)
		for v := range this.Neighbours[current] {
			_, found := unvisited[v]
			if found {
				distance := distances[current] + this.EdgesMap[fmt.Sprintf("%d-%d", current, v)]
				if distance < distances[v] {
					distances[v] = distance
				}
			}
		}
		current = next(unvisited, distances)
		if current == -1 {
			return -1
		}
	}

	return distances[node2]
}

func next(unvisited map[int]bool, distances map[int]int) int {
	node := -1
	min := int(1e7)
	for v := range unvisited {
		distance := distances[v]
		if distance < min {
			node = v
			min = distance
		}
	}

	return node
}

/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */
