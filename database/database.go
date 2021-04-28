package database

import (
	"errors"
	"fmt"

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

/*
	Selects rows from the given relation.
	The function signature reads like:
		SELECT * FROM from WHERE from.column condition(row, from.column)
*/
func Select(from Relation, column string, condition Condition) (Relation, error) {
	idx, err := util.IndexOf(column, from.Schema)
	if err != nil {
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
		idx, _ := util.IndexOf(column, r.Schema)
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
