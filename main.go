package main

import (
	"fmt"

	"github.com/zachtylr21/datalog-interpreter/database"
	"github.com/zachtylr21/datalog-interpreter/graph"
	"github.com/zachtylr21/datalog-interpreter/parser"
)

func main() {
	var program parser.DatalogProgram
	program.Run("test.txt")

	ruleGraph := program.RuleDependencies()
	fmt.Println(ruleGraph.String())

	sccs := graph.StronglyConnectedComponents(ruleGraph)

	fmt.Println(sccs)

	var db database.Database
	db.Create(program)

	C := db.Relations["C"]
	fmt.Println(C.Select("G", C.EqualsColumn("H")))
	fmt.Println(C.Select("H", C.Equals("'y'")))

	// fmt.Println(program.schemes)
	// fmt.Println(program.facts)
	// fmt.Println(program.rules)
	// fmt.Println(program.queries)
}
