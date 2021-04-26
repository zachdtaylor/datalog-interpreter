package database

import (
	"github.com/zachtylr21/datalog-interpreter/parser"
)

type Database struct {
	relations map[string]Relation
}

func (d *Database) Create(dp parser.DatalogProgram) {
	d.relations = make(map[string]Relation)

	for _, scheme := range dp.Schemes() {
		relationID := scheme.GetID()
		tuples := []Tuple{}
		for _, fact := range dp.Facts() {
			if fact.GetID() == relationID {
				tuple := fact.GetParameterValues()
				tuples = append(tuples, tuple)
			}
		}
		relation := Relation{scheme, tuples}
		d.relations[relation.Name()] = relation
	}
}

func (d *Database) Relations() map[string]Relation {
	return d.relations
}
