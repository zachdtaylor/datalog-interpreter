package graph

import (
	"strconv"

	"github.com/zachtylr21/datalog-interpreter/util"
)

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
func (g *Graph) DepthFirstSearch(node int) util.Stack {
	stack := util.Stack{}
	g.dfs(node, &stack)
	return stack
}

func (g *Graph) dfs(node int, stack *util.Stack) {
	if g.dependencyList[node].visited {
		return
	}
	g.dependencyList[node].visited = true
	dependencies := g.dependencyList[node].dependencies.Array()
	for _, node := range dependencies {
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
func (g *Graph) DFSForest() util.Stack {
	stack := util.Stack{}
	for node := range g.dependencyList {
		g.dfs(node, &stack)
	}
	return stack
}

func (g *Graph) String() string {
	graph := ""
	for id, node := range g.dependencyList {
		graph += strconv.Itoa(id) + ":"
		for _, dep := range node.dependencies.Array() {
			graph += strconv.Itoa(dep) + ","
		}
		graph += "\n"
	}
	return graph
}

func StronglyConnectedComponents(graph Graph) [][]int {
	transpose := Transpose(graph)
	finishOrder := graph.DFSForest()
	var sccs [][]int
	for range finishOrder.Values() {
		startNode := finishOrder.Pop()
		scc := transpose.DepthFirstSearch(startNode)
		if len(scc.Values()) != 0 {
			sccs = append(sccs, scc.Values())
		}
	}
	return sccs
}

func Transpose(graph Graph) Graph {
	var transpose Graph
	transpose.Init()
	for id, node := range graph.dependencyList {
		transpose.AddNode(id)
		for _, dep := range node.dependencies.Array() {
			transpose.AddDependency(dep, id)
		}
	}
	return transpose
}
