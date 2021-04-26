package main

import (
	"fmt"

	"github.com/zachtylr21/datalog-interpreter/database"
	"github.com/zachtylr21/datalog-interpreter/parser"
)

func main() {
	var program parser.DatalogProgram
	program.Init()
	program.Run("test.txt")

	var database database.Database
	database.Create(program)

	fmt.Println(database.Relations())

	// fmt.Println(program.schemes)
	// fmt.Println(program.facts)
	// fmt.Println(program.rules)
	// fmt.Println(program.queries)
}
