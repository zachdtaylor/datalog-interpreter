package graph

import "github.com/zachtylr21/datalog-interpreter/util"

type Node struct {
	value        int
	dependencies util.IntSet
	visited      bool
}

func (n *Node) AddDependency(value int) {
	n.dependencies.Add(value)
}
