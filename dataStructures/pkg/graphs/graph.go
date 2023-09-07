package graphs

import (
	"errors"
	"fmt"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) AddEdge(from, to int) error {
	fromVertex, err := g.getVertex(from)
	if err != nil {
		return err
	}

	toVertex, err := g.getVertex(to)
	if err != nil {
		return err
	}

	// Check if the edge already exists
	for _, v := range fromVertex.adjacent {
		if v == toVertex {
			return errors.New(fmt.Sprintf("Edge from %d to %d already exists", from, to))
		}
	}

	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)

	return nil
}

func (g *Graph) getVertex(k int) (*Vertex, error) {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i], nil
		}
	}
	return nil, errors.New(fmt.Sprintf("key %d not found", k))
}

func (g *Graph) AddVertex(k int) error {
	if g.contains(k) {
		return errors.New(fmt.Sprintf("key %d already exists", k))
	}
	g.vertices = append(g.vertices, &Vertex{key: k})
	return nil
}

func (g *Graph) contains(k int) bool {
	for _, v := range g.vertices {
		if v.key == k {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("%d -> ", v.key)
		for _, a := range v.adjacent {
			fmt.Printf("%d ", a.key)
		}
		fmt.Println()
	}
}

func (g *Graph) BFS(startKey int, applyFunc func(vertex *Vertex)) error {
	startVertex, err := g.getVertex(startKey)
	if err != nil {
		return err
	}

	visited := make(map[*Vertex]bool)
	queue := []*Vertex{startVertex}

	for len(queue) > 0 {
		vertex := queue[0]
		queue := queue[1:]
		if visited[vertex] {
			continue
		}

		visited[vertex] = true
		applyFunc(vertex)

		for _, neighbor := range vertex.adjacent {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
	return nil
}
