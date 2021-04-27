package database

import (
	"errors"
	"fmt"

	"github.com/zachtylr21/datalog-interpreter/util"
)

type Tuple []string

type Condition func(Tuple, string) bool

type Relation struct {
	Name   string
	Schema []string
	Tuples []Tuple
}

func (r Relation) Equals(value string) Condition {
	condition := func(row Tuple, compareTo string) bool {
		return compareTo == value
	}
	return condition
}

func (r Relation) EqualsColumn(column string) Condition {
	idx, _ := util.IndexOf(column, r.Schema)
	condition := func(row Tuple, compareTo string) bool {
		return row[idx] == compareTo
	}
	return condition
}

func (r Relation) Select(column string, condition Condition) (Relation, error) {
	idx, err := util.IndexOf(column, r.Schema)
	if err != nil {
		relation := Relation{Name: r.Name, Schema: r.Schema}
		return relation, errors.New(fmt.Sprintf("column %s does not exist", column))
	}
	tuples := []Tuple{}
	for _, tuple := range r.Tuples {
		if condition(tuple, tuple[idx]) {
			tuples = append(tuples, tuple)
		}
	}
	relation := Relation{r.Name, r.Schema, tuples}
	return relation, nil
}
