package database

type Tuple []string

type Relation struct {
	Name   string
	Schema []string
	Tuples []Tuple
}
