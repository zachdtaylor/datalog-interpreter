package database

import "github.com/zachtylr21/datalog-interpreter/parser"

type Tuple []string

type Relation struct {
	scheme parser.Scheme
	tuples []Tuple
}

func (r *Relation) Name() string {
	return r.scheme.GetID()
}
