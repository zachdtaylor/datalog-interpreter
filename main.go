package main

import (
	"fmt"

	. "github.com/zachtylr21/datalog-interpreter/database"
	"github.com/zachtylr21/datalog-interpreter/parser"
)

func main() {
	var program parser.DatalogProgram
	program.Run("test.txt")

	// ruleGraph := program.RuleDependencies()

	// sccs := graph.StronglyConnectedComponents(ruleGraph)

	var db Database
	db.Create(program)

	C := db.Relations["C"]
	s1, _ := Select(C, "G", EqualsColumn("H"))
	s2, _ := Select(C, "H", Equals("'y'"))
	fmt.Println(s1.Tuples)
	fmt.Println(s2.Tuples)

	// fmt.Println(program.schemes)
	// fmt.Println(program.facts)
	// fmt.Println(program.rules)
	// fmt.Println(program.queries)
}
