package main

import (
	"github.com/zachtylr21/datalog-interpreter/parser"
)

func main() {
	var program parser.DatalogProgram
	program.Init()
	program.Run("test.txt")

	// fmt.Println(program.schemes)
	// fmt.Println(program.facts)
	// fmt.Println(program.rules)
	// fmt.Println(program.queries)
}
