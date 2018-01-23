package main

import "fmt"

func main() {
	vertices := []Vertex{
		Vertex{id: 1, edges: []int{2, 3}},
		Vertex{id: 2, edges: []int{4, 5}},
		Vertex{id: 3, edges: []int{6, 7}},
		Vertex{id: 4},
		Vertex{id: 5, edges: []int{7, 11}},
		Vertex{id: 6, edges: []int{10, 11}},
		Vertex{id: 7, edges: []int{8, 9}},
		Vertex{id: 8},
		Vertex{id: 9},
		Vertex{id: 10},
		Vertex{id: 11}}
	paths := [][]int{}
	currentPath := []int{}
	visited := []int{}
	getShortestPath(vertices, 1, 9, &paths, currentPath, visited)
	shortestPath := selectShortestPath(paths)
	fmt.Println(vertices, paths, shortestPath)
}

type Vertex struct {
	id    int
	edges []int
}

func getShortestPath(vertices []Vertex, origin int, destination int, paths *[][]int, currentPath []int, visited []int) {
	v := selectVertexById(vertices, origin)
	search(vertices, v, destination, paths, currentPath, visited)
}

func selectVertexById(vertices []Vertex, id int) Vertex {
	for _, v := range vertices {
		if v.id == id {
			return v
		}
	}
	return Vertex{}
}

func search(vertices []Vertex, v Vertex, destination int, paths *[][]int, currentPath []int, visited []int) {
	visited = append(visited, v.id)
	currentPath = append(currentPath, v.id)
	vIndex := len(currentPath) - 1
	for _, e := range v.edges {
		if e == destination {
			currentPath = append(currentPath, destination)
			appendPointer(paths, currentPath)
			currentPath = remove(currentPath, len(currentPath)-1) // last index
		} else {
			if contains(visited, destination) == false {
				w := selectVertexById(vertices, e)
				search(vertices, w, destination, paths, currentPath, visited)
			}
		}
	}
	currentPath = remove(currentPath, vIndex)
}

func selectShortestPath(paths [][]int) []int {
	shortest := []int{}
	if len(paths) > 1 {
		shortest = paths[0]
	}
	for _, p := range paths {
		if len(p) < len(shortest) {
			shortest = p
		}
	}
	return shortest
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func appendPointer(sp *[][]int, i []int) {
	s := *sp
	*sp = append(s, i)
}
