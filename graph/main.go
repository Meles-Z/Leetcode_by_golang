package main

import (
	"fmt"
)

// graph structure

type Graph struct {
	vertices []*Vertex
}

// vertex structure
type Vertex struct {
	Key       int
	adjencent []*Vertex
}

func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v are aready exist", k)
		fmt.Println(err.Error())
	} else {

		g.vertices = append(g.vertices, &Vertex{Key: k})
	}
}

func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.Key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *Graph) print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v: ", v.Key)
		for _, v := range v.adjencent {
			fmt.Printf("%v", v.Key)
		}
	}
	fmt.Println()
}

// contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.Key {
			return true
		}
	}
	return false

}

// add edge
func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check errors
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjencent, to) {
		err := fmt.Errorf("existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {

		//add edge
		fromVertex.adjencent = append(fromVertex.adjencent, toVertex)
	}

}

func main() {
	test := Graph{}
	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}
	test.AddVertex(9)
	test.AddVertex(10)
	test.AddEdge(1, 2)
	test.AddEdge(1, 2)
	test.print()
}
