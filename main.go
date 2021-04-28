package main

import (
	"fmt"

	. "github.com/zachtylr21/datalog-interpreter/database"
	"github.com/zachtylr21/datalog-interpreter/parser"
)

func main() {
	var program parser.DatalogProgram
	program.Run("test.txt")

	var db Database
	db.Create(program)

	B := db.Relations["B"]
	A := db.Relations["A"]
	C := db.Relations["C"]

	fmt.Println(Project(B, []string{"W"}))

	R := Join(A, C)
	fmt.Println(R)
}
