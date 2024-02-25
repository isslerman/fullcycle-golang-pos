package main

import "fmt"

// Graph represents an adjacency list graph
type Graph struct {
	vertices []*Vertex
}

// Vertex represents a graph vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added. Key already exist", k)
		fmt.Println(err)
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check for errors
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Error: invalid edge %v - %v", from, to)
		fmt.Println(err)
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Vertex %v not added. Key already exist", from)
		fmt.Println(err)
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

func (g *Graph) getVertex(k int) *Vertex {
	for _, v := range g.vertices {
		if k == v.key {
			return v
		}
	}
	err := fmt.Errorf("Error: Vertex %v not found", k)
	fmt.Println(err)
	return nil
}

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v :", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" -> %v", v.key)
		}
	}
}

func main() {
	test := &Graph{}

	test.AddVertex(1)
	test.AddVertex(2)
	test.AddVertex(3)
	test.AddVertex(4)
	test.AddEdge(1, 2)

	test.AddEdge(1, 6)
	test.Print()
}
