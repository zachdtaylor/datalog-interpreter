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

/*
  Performs depth first search from the given starting node.
	Returns a stack which holds the nodes in order of their finish times, from
	highest (top of the stack) to lowest.
*/
func (g *Graph) DepthFirstSearch(node int) Stack {
	stack := Stack{}
	g.dfs(node, &stack)
	return stack
}

func (g *Graph) dfs(node int, stack *Stack) {
	if g.dependencyList[node].visited {
		return
	}
	g.dependencyList[node].visited = true
	for node := range g.dependencyList[node].dependencies.values {
		if !g.dependencyList[node].visited {
			g.dfs(node, stack)
		}
	}
	stack.Push(node)
}

/*
  Performs depth first search from a random starting node, then continues
	to do so until all nodes have been visited.
	Returns a stack which holds the nodes in order of their finish times, from
	highest (top of the stack) to lowest.
*/
func (g *Graph) DFSForest() Stack {
	stack := Stack{}
	for node := range g.dependencyList {
		g.dfs(node, &stack)
	}
	return stack
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
