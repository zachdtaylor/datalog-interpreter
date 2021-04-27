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

/*
  Computes the strongly connected components (SCCs) of the given graph.
	SCCs can be computed using Kosaraju's algorithm:
	  1. Perform depth first search forest on the graph and keep track of
		   the finish times of each node.
		2. Compute the transpose graph
		3. For each node n in descending finish time order:
		   - If n has been visited in the transpose graph, continue to next loop.
		   - Perform depth first search on the transpose graph starting from node n,
			   keeping track of which nodes were visited.
			 - The set of nodes visited is an SCC. Store the SCC, and continue.
*/
func StronglyConnectedComponents(graph Graph) [][]int {
	finishOrder := graph.DFSForest()
	transpose := Transpose(graph)
	var sccs [][]int
	for range finishOrder.Values() {
		startNode := finishOrder.Pop()
		if transpose.dependencyList[startNode].visited {
			continue
		}
		scc := transpose.DepthFirstSearch(startNode)
		sccs = append(sccs, scc.Values())
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
