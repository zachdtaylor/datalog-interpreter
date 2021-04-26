package main

import (
	"fmt"

	"github.com/zachtylr21/datalog-interpreter/database"
	"github.com/zachtylr21/datalog-interpreter/parser"
)

func main() {
	var program parser.DatalogProgram
	program.Run("test.txt")

	graph, revGraph := program.RuleDependencies()
	fmt.Println(graph.String())
	fmt.Println(revGraph.String())

	var database database.Database
	database.Create(program)

	fmt.Println(database.Relations())

	// fmt.Println(program.schemes)
	// fmt.Println(program.facts)
	// fmt.Println(program.rules)
	// fmt.Println(program.queries)
}
