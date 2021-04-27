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

	SCCs can be computed using Kosaraju's algorithm, which is what we use here.
	This algorithm works because of the following fact:
	  > If we do DFS forest on a graph, the finish time of a node that connects to other
		> SCCs will always be greater than the finish time of nodes in the other SCC.
	To illustrate this fact, say we have the following graph:
	  0: 2, 3
		1: 0
		2: 1
		3: 4
		4:
	The SCCs are G1: [0, 1, 2], G2: [3], and G3: [4], and observe that 0 connects from
	G1 to G2 and 3 connects from G2 to G3.
	Say we start DFS forest at node 3. The traversal path is 3 -> 4, so the nodes in order
	of finish time (lowest to highest) is [4, 3]. Then say we do another pass starting at
	node 1. A traversal path is 1 -> 0 -> 2, so the finish time order is [2, 0, 1]. The
	complete finish time order is [4, 3, 2, 0, 1]. Notice that 0 has a greater finish time
	than 3 and 3 has a greater finish time than 4.

	Now to understand why the algorithm works, consider the reverse graph:
	  0: 1
		1: 2
		2: 0
		3: 0
		4: 3
	Notice that [0, 1, 2] is a sink in this graph, and if we remove those nodes, then
	[3] is a sink in the remaining graph, and if we remove 3 as well, then [4] is a sink
	in the remaining graph. So, getting the SCCs is as simple as doing DFS 3 times on
	the reverse graph starting with either 0, 1, or 2 first, then starting with 3, then
	starting with 4. The complete finish time order from the DFS forest on the original
	graph is the exact order we need.
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
