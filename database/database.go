package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/zachtylr21/datalog-interpreter/parser"
	"github.com/zachtylr21/datalog-interpreter/util"
)

type Condition func(Relation, Tuple, string) bool

type Database struct {
	Relations map[string]Relation
}

func (d *Database) Create(dp parser.DatalogProgram) {
	d.Relations = make(map[string]Relation)

	for _, scheme := range dp.Schemes() {
		relationID := scheme.GetID()
		tuples := []Tuple{}
		for _, fact := range dp.Facts() {
			if fact.GetID() == relationID {
				tuple := fact.GetParameterValues()
				tuples = append(tuples, tuple)
			}
		}
		relation := Relation{scheme.GetID(), scheme.GetParameterValues(), tuples}
		d.Relations[relation.Name] = relation
	}
}

func Rename(r Relation, schema Schema) Relation {
	return Relation{r.Name, schema, r.Tuples}
}

func Project(r Relation, cols []string) Relation {
	var tuples []Tuple
	for _, t := range r.Tuples {
		var tuple Tuple
		for _, col := range cols {
			tuple = append(tuple, t[util.IndexOf(col, r.Schema)])
		}
		tuples = append(tuples, tuple)
	}
	return Relation{r.Name, cols, tuples}
}

func Join(r1, r2 Relation) Relation {
	schema := joinSchemas(r1.Schema, r2.Schema)
	var tuples []Tuple
	for _, tup1 := range r1.Tuples {
		for _, tup2 := range r2.Tuples {
			tup, err := joinTuples(r1.Schema, r2.Schema, tup1, tup2)
			if err == nil {
				tuples = append(tuples, tup)
			}
		}
	}
	return Relation{r1.Name, schema, tuples}
}

func joinTuples(schema1, schema2 Schema, tuple1, tuple2 Tuple) (Tuple, error) {
	tuple := tuple1
	for idx2, col := range schema2 {
		idx1 := util.IndexOf(col, schema1)
		if idx1 == -1 {
			tuple = append(tuple, tuple2[idx2])
		} else if tuple1[idx1] != tuple2[idx2] {
			return nil, errors.New("Tuples are not joinable")
		}
	}
	return tuple, nil
}

func joinSchemas(schema1, schema2 Schema) Schema {
	for _, col := range schema2 {
		idx := util.IndexOf(col, schema1)
		if idx == -1 { // col is not in schema1
			schema1 = append(schema1, col)
		}
	}
	return schema1
}

/*
	Selects rows from the given relation.
	The function signature reads like:
		SELECT * FROM from WHERE from.column condition
*/
func Select(from Relation, column string, condition Condition) (Relation, error) {
	idx := util.IndexOf(column, from.Schema)
	if idx == -1 {
		relation := Relation{Name: from.Name, Schema: from.Schema}
		return relation, errors.New(fmt.Sprintf("column %s does not exist", column))
	}
	tuples := []Tuple{}
	for _, tuple := range from.Tuples {
		if condition(from, tuple, tuple[idx]) {
			tuples = append(tuples, tuple)
		}
	}
	relation := Relation{from.Name, from.Schema, tuples}
	return relation, nil
}

func Equals(value string) Condition {
	condition := func(r Relation, row Tuple, compareTo string) bool {
		return compareTo == value
	}
	return condition
}

/*
  This method is ineffecient. If performance is an issue, consider
	using EqualsColumnIdx
*/
func EqualsColumn(column string) Condition {
	condition := func(r Relation, row Tuple, compareTo string) bool {
		idx := util.IndexOf(column, r.Schema)
		if idx == -1 {
			log.Fatal(errors.New(fmt.Sprintf("Column %s does not exist on relation %s", column, r.Name)))
		}
		return row[idx] == compareTo
	}
	return condition
}

/*
  A more time-effecient version of EqualsColumn that does not need to look up
	the index of the column for each row, in exchange for the user computing the
	index beforehand.
	Example:
		C := db.Relations["C"]
		Select(C, "G", EqualsColumn("H")) // <- this is slower
		idx := util.IndexOf("H", C.Schema)
		Select(C, "G", EqualsColumnIdx(idx)) // <- this is faster!
*/
func EqualsColumnIdx(column int) Condition {
	condition := func(r Relation, row Tuple, compareTo string) bool {
		return row[column] == compareTo
	}
	return condition
}
