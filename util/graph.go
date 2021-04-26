package util

import (
	"strconv"
)

type Node struct {
	value        int
	dependencies IntSet
	visited      bool
}

func (n *Node) AddDependency(value int) {
	n.dependencies.Add(value)
}

type Graph struct {
	dependencyList map[int]*Node
}

func (g *Graph) Init() {
	g.dependencyList = make(map[int]*Node)
}

func (g *Graph) AddNode(value int) {
	if _, ok := g.dependencyList[value]; !ok {
		g.dependencyList[value] = &Node{value: value}
	}
}

func (g *Graph) AddDependency(value, dep int) {
	if _, ok := g.dependencyList[value]; !ok {
		g.AddNode(value)
	}
	g.dependencyList[value].AddDependency(dep)
}

func (g *Graph) String() string {
	graph := ""
	for _, node := range g.dependencyList {
		graph += strconv.Itoa(node.value) + ":"
		for _, dep := range node.dependencies.Array() {
			graph += strconv.Itoa(dep) + ","
		}
		graph += "\n"
	}
	return graph
}
